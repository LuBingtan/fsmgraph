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
			if nil == edge.NextEdge {
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

func (g *Graph) AddEdge(src, dst *Vertex) error {
	// if the vertex is new, set it in the vertex map
	// set start point
	i, ok := g.vertexMap[src.Id]
	if !ok || i != src.Index {
		src.Outdegree = 0
		src.Index = len(g.Vertexlist)

		g.Vertexlist = append(g.Vertexlist, src)
		g.vertexMap[src.Id] = src.Index
	}
	g.Vertexlist[g.vertexMap[src.Id]].Outdegree++

	// set end point
	i, ok = g.vertexMap[dst.Id]
	if !ok || i != dst.Index {
		dst.Indegree = 0
		dst.Index = len(g.Vertexlist)

		g.Vertexlist = append(g.Vertexlist, dst)
		g.vertexMap[dst.Id] = dst.Index
	}
	g.Vertexlist[g.vertexMap[dst.Id]].Indegree++

	// set vertex edge
	g.vertexMap[src.Id].Adjoin(g.vertexMap[dst.Id])

	return nil
}

func (g *Graph) TopoSort() (sortIndexList []int, err error) {
	indgreeList := []int{}

	// put all indegree in list
	for _, v := range g.Vertexlist {
		indgreeList = append(indgreeList, v.Indegree)
	}

	// find 0 indgree index
	for i, d := range indgreeList {
		if 0 == d {
			sortIndexList = append(sortIndexList, i)
		}
	}

	// 0 indgree adjoin vertex degree minus 1
	// TODO: need to be more elegant
	for i := 0; i < len(sortIndexList); i++ {
		// get index
		index := sortIndexList[i]

		v := g.Vertexlist[index]
		edge := v.OutEdge
		for {
			if nil == edge {
				break
			}

			indgreeList[edge.AdjIndex]--
			if 0 == indgreeList[edge.AdjIndex] {
				sortIndexList = append(sortIndexList, edge.AdjIndex)
			}

			edge = edge.NextEdge
		}
	}

	//sort.Sort(g.Vertexlist)
	return sortIndexList, nil
}

func (g *Graph) BFS() (sortOut []*Vertex) {
	for _, v := range g.Vertexlist {

	}
	return
}

func (g *Graph) DFS() (sortOut []*Vertex) {
	return
}
