package graph

import (
	"container/list"
	"fmt"
	//"sort"
)

var l = list.New()

type GraphType string

type GraphInterface interface {
	AddEdge(src, dst *Vertex) error
	TopoSort() ([]*Vertex, error)
	BFS() []*Vertex
	DFS() []*Vertex
	Type() GraphType
}

type ExecutorFunc func(...interface{}) (interface{}, error)

type VertexState string

/// [Define]
// structure for vertex
type VertexSlice []*Vertex

func (s VertexSlice) Len() int {
	return len(s)
}

func (s VertexSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s VertexSlice) Less(i, j int) bool {
	return s[i].Index < s[j].Index
}

type Vertex struct {
	// meta data
	Id   string
	Data interface{}
	// graph data
	Index     int
	InEdge    *Edge
	OutEdge   *Edge
	Indegree  int
	Outdegree int
	// state data
	State    VertexState
	Executor ExecutorFunc
}

func NewVertex(id string) *Vertex {
	return &Vertex{Id: id}
}

func (v *Vertex) Adjoin(dst *Vertex) error {
	if dst.Id == v.Id {
		return fmt.Errorf("vertex id[%s] conflict", v.Id)
	}

	if nil == v.OutEdge {
		v.OutEdge = &Edge{
			AdjIndex: dst.Index,
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
			AdjIndex: dst.Index,
		}
	}

	return nil
}

/// [Define]
// structure for Edge

type EdgeState string

type Edge struct {
	// meta data
	Weight int
	// graph data
	AdjIndex int
	NextEdge *Edge
	// state data
	State EdgeState
}

type Graph struct {
	vertexMap  map[string]*Vertex
	Vertexlist VertexSlice
}

func NewGraph() Graph {
	return Graph{
		vertexMap: make(map[string]*Vertex),
	}
}

func (g *Graph) AddEdge(src, dst *Vertex) error {
	// if the vertex is new, set it in the vertex map
	// set start point
	v, ok := g.vertexMap[src.Id]
	if !ok || v != src {
		g.vertexMap[src.Id] = src
		g.vertexMap[src.Id].Outdegree = 0
		g.vertexMap[src.Id].Index = len(g.Vertexlist)
		g.Vertexlist = append(g.Vertexlist, g.vertexMap[src.Id])
	}
	g.vertexMap[src.Id].Outdegree++

	// set end point
	v, ok = g.vertexMap[dst.Id]
	if !ok || v != dst {
		g.vertexMap[dst.Id] = dst
		g.vertexMap[dst.Id].Indegree = 0
		g.vertexMap[dst.Id].Index = len(g.Vertexlist)
		g.Vertexlist = append(g.Vertexlist, g.vertexMap[dst.Id])
	}
	g.vertexMap[dst.Id].Indegree++

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

	return
}

func (g *Graph) DFS() (sortOut []*Vertex) {
	return
}
