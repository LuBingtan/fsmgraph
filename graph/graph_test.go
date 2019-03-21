package graph

import (
	"testing"
)

func GraphPrint(t *testing.T, g *Graph) {
	t.Logf("Graph node number:%d\n", g.Size())
	g.TopoSort()
	//t.Log("=======sort=======")
	for _, v := range g.VertexList() {
		t.Log("=========Node=========:")
		t.Log("id:", v.Id)
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
	GraphPrint(t, g)

}

func Test4Graph4Delete(t *testing.T) {
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

	//GraphPrint(t, g)

	g.DeleteVertex("3")
	GraphPrint(t, g)
}