package graph

import (
	"fmt"
	simpleSt "fsmgraph-lib/simplestructure"
)

/*****************************************  graph interface  *****************************************/

type GraphType string

type GraphInterface interface {
	// Create
	InsertVertex(string, interface{}) *Vertex
	InsertEdge(src, dst *Vertex) error

	// Retrieve & Update

	TopoSort() ([]*Vertex, error)
	BFS() []*Vertex
	DFS() []*Vertex
	Type() GraphType
}

/*****************************************  vertex  *****************************************/

type ExecutorFunc func(...interface{}) (interface{}, error)

type VertexState string

/// [Define]
// structure for vertex
type Vertex struct {
	// meta data
	Id   string
	Data interface{}
	// graph data
	InEdge    *Edge
	OutEdge   *Edge
	Indegree  int
	Outdegree int
	// state data
	State    VertexState
	Executor ExecutorFunc
}

func NewVertex(id string, data interface{}) *Vertex {
	return &Vertex{
		Id:        id,
		Data:      data,
		InEdge:    nil,
		OutEdge:   nil,
		Indegree:  0,
		Outdegree: 0,
	}
}

func (v *Vertex) AddNext(next *Vertex) {
	if nil == v.OutEdge {
		v.OutEdge = &Edge{
			AdjVertex: next,
		}
	} else {
		edge := v.OutEdge
		for {
			if next == edge.AdjVertex {
				return
			}
			if nil == edge.NextEdge {
				break
			}
			edge = edge.NextEdge
		}
		edge.NextEdge = &Edge{
			AdjVertex: next,
		}
	}
	v.Outdegree++
}

func (v *Vertex) AddPrev(prev *Vertex) {
	if nil == v.InEdge {
		v.InEdge = &Edge{
			AdjVertex: prev,
		}
	} else {
		edge := v.InEdge
		for {
			if nil == edge.NextEdge || prev == edge.AdjVertex {
				break
			}
			edge = edge.NextEdge
		}
		edge.NextEdge = &Edge{
			AdjVertex: prev,
		}
	}
	v.Indegree++
}

func (v *Vertex) RemoveNext() {
	edge := v.OutEdge

	for {
		if nil == edge {
			break
		}

		edge.AdjVertex.InEdge.AdjVertex = nil

		edge = edge.NextEdge
	}

	v.OutEdge = nil
}

func (v *Vertex) RemovePrev() {
	edge := v.InEdge

	for {
		if nil == edge {
			break
		}

		edge.AdjVertex.OutEdge.AdjVertex = nil

		edge = edge.NextEdge
	}

	v.OutEdge = nil
}

/*****************************************  edge  *****************************************/

/// [Define]
// structure for Edge
type EdgeState string

type Edge struct {
	// meta data
	Weight int
	// graph data
	AdjVertex *Vertex
	NextEdge  *Edge
	// state data
	State EdgeState
}

/*****************************************  garph  *****************************************/

/// [Define]
// structure for graph
type Graph struct {
	// this map is used to ensure uniqueness of vertex
	vertexMap    map[string]int
	vertexVector *simpleSt.SimpleVector
}

func NewGraph() Graph {
	return Graph{
		vertexMap:    make(map[string]int),
		vertexVector: simpleSt.NewSimpleVector(),
	}
}

func (g *Graph)Size() int {
	return g.vertexVector.Len()
}

func (g *Graph) InsertVertex(id string, data interface{}) *Vertex {
	v := NewVertex(id, data)

	index, ok := g.vertexMap[id]
	if !ok {
		// insert new vertex
		g.vertexVector.Pushback(v)
		g.vertexMap[v.Id] = g.vertexVector.Len() - 1
	} else {
		// init graph relation data
		v = g.vertexVector.At(index).(*Vertex)
		v.Data = data
	}

	return v
}

func (g *Graph) GetVertex(id string) *Vertex {
	index, ok := g.vertexMap[id]
	if !ok {
		return nil
	}

	return g.vertexVector.At(index).(*Vertex)
}

func (g *Graph) GetVertexByIndex(index int) *Vertex {
	v := g.vertexVector.At(index)
	if v == nil {
		return nil
	}
	return v.(*Vertex)
}

func (g *Graph) DeleteVertex(id string) error {
	index, ok := g.vertexMap[id]
	if !ok {
		return fmt.Errorf("unknown vertext[id:%s].", id)
	}

	// delete map
	delete(g.vertexMap, id)

	// delete vertex vector
	v := g.vertexVector.Remove(index).(*Vertex)

	// remove adjoin vertex
	v.RemoveNext()
	v.RemovePrev()

	return nil
}

func (g *Graph) AddEdge(src, dst string) error {
	// if the vertex is new, set it in the vertex map
	srcVertex := g.GetVertex(src)
	if srcVertex == nil {
		return fmt.Errorf("unknown vertex[id:%s]", src)
	}

	dstVertex := g.GetVertex(dst)
	if dstVertex == nil {
		return fmt.Errorf("unknown vertex[id:%s]", dst)
	}

	srcVertex.AddNext(dstVertex)
	dstVertex.AddPrev(srcVertex)

	return nil
}

func (g *Graph) VertexList(indexList ...int) (vertexList []*Vertex) {
	if len(indexList) != 0 {
		for _, index := range indexList {
			vertexList = append(vertexList, g.GetVertexByIndex(index))
		}
		return vertexList
	}
	it := g.vertexVector.Iterator()
	for {
		next := it.Next()
		if next == nil {
			break
		}
		vertexList = append(vertexList, next.(*Vertex))
	}
	return vertexList
}

func (g *Graph) TopoSort() (sortVertexList []*Vertex, err error) {
	indgreeList := []int{}
	indexQueue := simpleSt.NewSimpleQueue()

	// put all indegree in list
	it := g.vertexVector.Iterator()
	for {
		next := it.Next()
		if next == nil {
			break
		}

		v := next.(*Vertex)
		indgreeList = append(indgreeList, v.Indegree)
	}

	// find 0 indgree index
	for i, d := range indgreeList {
		if 0 == d {
			indexQueue.Pushback(i)
		}
	}

	// 0 indgree adjoin vertex degree minus 1
	// TODO: need to be more elegant
	for {
		// get index
		p := indexQueue.Popfront()
		if p == nil {
			break
		}
		index := p.(int)
		sortVertexList = append(sortVertexList, g.GetVertexByIndex(index))
		// ge vertex
		v := g.vertexVector.At(index).(*Vertex)
		edge := v.OutEdge
		for {
			if nil == edge {
				break
			}

			adjIndex := g.vertexMap[edge.AdjVertex.Id]
			indgreeList[adjIndex]--
			if 0 == indgreeList[adjIndex] {
				indexQueue.Pushback(adjIndex)
			}

			edge = edge.NextEdge
		}
	}

	//sort.Sort(g.Vertexlist)
	return sortVertexList, nil
}

func (g *Graph) BFS() (sortOut []*Vertex) {
	return
}

func (g *Graph) DFS() (sortOut []*Vertex) {
	return
}
