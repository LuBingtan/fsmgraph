package graph

import (
	"fmt"
)

type GraphType string

type GraphInterface interface {
	AddEdge(src, dst *Vertex) error
	TopoSort() ([]*Vertex, error)
	BFS() []*Vertex
	DFS() []*Vertex
	Type() GraphType
}

type ExecutorFunc func(interface{}) (interface{}, error)

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
	return s[i].Indegree < s[j].Indegree
}

type Vertex struct {
	// meta data
	Id        string
	Data      interface{}
	// graph data
	Index     int
	FirstEdge *Edge
	Indegree  int
	Outdegree int
	// state data
	State     VertexState
	Executor  ExecutorFunc
}

func NewVertex(id string) *Vertex {
	return &Vertex{Id: id}
}

func (v *Vertex) Adjoin(dst *Vertex) error {
	if dst.Id == v.Id {
		return fmt.Errorf("vertex id[%s] conflict", v.Id)
	}

	if nil == v.FirstEdge {
		v.FirstEdge = &Edge{
			AdjIndex: dst.Index,
		}
	} else {
		edge := v.FirstEdge
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
	Weight   int
	// graph data
	AdjIndex int
	NextEdge *Edge
	// state data
	State    EdgeState
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

func (g *Graph) TopoSort() (vertexList []Vertex, err error) {

	

	return vertexList, nil
}

func (g *Graph) BFS() (sortOut []*Vertex) {

	return
}

func (g *Graph) DFS() (sortOut []*Vertex) {
	return
}
