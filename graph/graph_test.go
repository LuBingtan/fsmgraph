package graph

import (
	"testing"
)

func Test4Graph4Init(t *testing.T) {
	t.Logf("testing for graph initialize start.\n")
	g := NewGraph()
	g.InsertVertex("2", 2)
	g.InsertVertex("3", 3)
	g.InsertVertex("0", 0)
	g.InsertVertex("1", 1)

	g.AddEdge("1", "2")
	g.AddEdge("2", "3")
	g.AddEdge("0", "1")
	g.AddEdge("0", "2")
	t.Logf("Graph node number:%d\n", g.Size())
	s, _ := g.TopoSort()
	t.Log("=======sort=======")
	for m, v := range s {
		t.Log("=========Node=========:")
		t.Log("Index:", m)
		t.Log("Indegree:", v.Indegree)
		t.Log("Outdegree:", v.Outdegree)
		edge := v.OutEdge
		for {
			if nil == edge {
				break
			}
			t.Logf("edge:%v,", edge.AdjVertex)
			edge = edge.NextEdge
		}
	}

}

func Test4Graph4BFS(t *testing.T) {
	t.Logf("testing for graph tranverse start.\n")
	a := []int{}
	t.Log(len(a))
}
