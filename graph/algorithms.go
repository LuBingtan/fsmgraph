package graph
import simpleSt "fsmgraph-lib/simplestructure"
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

	return sortVertexList, nil
}

func (g *Graph) BFS(startPointId string) (sortVertexList []*Vertex) {
	
	return
}

func (g *Graph) DFS() (sortOut []*Vertex) {
	return
}
