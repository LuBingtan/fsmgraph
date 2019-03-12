package graph

import (
	"testing"
)

func Test4Graph4Init(t *testing.T) {
	t.Logf("testing for graph initialize start.")
	g := NewGraph()
	v1 := NewVertex(0)
	v2 := NewVertex(1)
	v3 := NewVertex(2)
	v4 := NewVertex(3)
	g.AddEdge(v3, v4)
	g.AddEdge(v1, v2)
	g.AddEdge(v1, v3)
	g.AddEdge(v2, v3)
	g.SortId()
	t.Logf("Graph node number:%d", len(g.Vertexlist))
	for _, v := range g.Vertexlist {
		t.Log("=========Node=========")
		t.Log("id:", v.Id)
		t.Log("Indegree:", v.Indegree)
		t.Log("Outdegree:", v.Outdegree)
		edge := v.FirstEdge
		for {
			if nil == edge {
				break
			}
			t.Logf("edge-%d,", edge.VertexId)
			edge = edge.NextEdge
		}
	}

}

func Test4Graph4BFS(t *testing.T) {
	t.Logf("testing for graph tranverse start.")
	a := []int{}
	t.Log(len(a))
}
