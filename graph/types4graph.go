package graph

import (
	"fmt"
	"sort"
)

type GraphType string

type GraphInterface interface {
	AddEdge(src, dst *Vertex) error
	Sort() ([]*Vertex, error)
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
	return s[i].Id < s[j].Id
}

type Vertex struct {
	Id        int
	Data      interface{}
	FirstEdge *Edge
	Indegree  int
	Outdegree int
	State     VertexState
	Executor  ExecutorFunc
}

func NewVertex(id int) *Vertex {
	return &Vertex{Id: id}
}

func (v *Vertex) Adjoin(dst *Vertex) error {
	if dst.Id == v.Id {
		return fmt.Errorf("vertex id[%d] conflict", v.Id)
	}

	if nil == v.FirstEdge {
		v.FirstEdge = &Edge{
			VertexId: dst.Id,
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
			VertexId: dst.Id,
		}
	}

	return nil
}

/// [Define]
// structure for Edge

type EdgeState string

type Edge struct {
	State    EdgeState
	Weight   int
	VertexId int
	NextEdge *Edge
}

type Graph struct {
	vertexMap  map[int]*Vertex
	Vertexlist VertexSlice
}

func NewGraph() Graph {
	return Graph{
		vertexMap: make(map[int]*Vertex),
	}
}

func (g *Graph) AddEdge(src, dst *Vertex) error {
	// if the vertex is new, set it in the vertex map
	// set start point
	v, ok := g.vertexMap[src.Id]
	if !ok || v != src {
		g.vertexMap[src.Id] = src
		g.vertexMap[src.Id].Outdegree = 0
		g.Vertexlist = append(g.Vertexlist, g.vertexMap[src.Id])
	}
	g.vertexMap[src.Id].Outdegree++

	// set end point
	v, ok = g.vertexMap[dst.Id]
	if !ok || v != dst {
		g.vertexMap[dst.Id] = dst
		g.vertexMap[dst.Id].Indegree = 0
		g.Vertexlist = append(g.Vertexlist, g.vertexMap[dst.Id])
	}
	g.vertexMap[dst.Id].Indegree++

	// set vertex edge
	g.vertexMap[src.Id].Adjoin(g.vertexMap[dst.Id])

	return nil
}

func (g *Graph) SortId() error {
	sort.Sort(g.Vertexlist)
	return nil
}

func (g *Graph) SortTopo() ([]*Vertex, error) {
	vertexList := []*Vertex{}
	return vertexList, nil
}

func (g *Graph) BFS() (sortOut []*Vertex) {

	return
}

func (g *Graph) DFS() (sortOut []*Vertex) {
	return
}
