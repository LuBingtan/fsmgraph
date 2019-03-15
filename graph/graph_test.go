package graph

import (
	"testing"
)

func Test4Graph4Init(t *testing.T) {
	t.Logf("testing for graph initialize start.\n")
	g := NewGraph()
	v1 := NewVertex("name0")
	v2 := NewVertex("name1")
	v3 := NewVertex("name2")
	v4 := NewVertex("name3")
	g.AddEdge(v3, v4)
	g.AddEdge(v1, v2)
	g.AddEdge(v1, v3)
	g.AddEdge(v2, v3)
	t.Logf("Graph node number:%d\n", len(g.Vertexlist))
	s, _ := g.TopoSort()
	t.Log("=======sort=======")
	for _, m := range s {
		t.Log(m)
	}
	for _, v := range g.Vertexlist {
		t.Log("=========Node=========")
		t.Log("Index:", v.Index)
		t.Log("Indegree:", v.Indegree)
		t.Log("Outdegree:", v.Outdegree)
		edge := v.OutEdge
		for {
			if nil == edge {
				break
			}
			t.Logf("edge:%d,", edge.AdjIndex)
			edge = edge.NextEdge
		}
		t.Logf("\n")
	}

}

func Test4Graph4BFS(t *testing.T) {
	t.Logf("testing for graph tranverse start.\n")
	a := []int{}
	t.Log(len(a))
}
