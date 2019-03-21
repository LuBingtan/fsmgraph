package graph

import (
	"fmt"
	simpleSt "fsmgraph-lib/simplestructure"
)

/*****************************************  garph  *****************************************/

/// [Define]
// structure for graph
type Graph struct {
	// this map is used to ensure uniqueness of vertex
	vertexMap    map[string]int
	vertexVector *simpleSt.SimpleVector
}

func NewGraph() *Graph {
	return &Graph{
		vertexMap:    make(map[string]int),
		vertexVector: simpleSt.NewSimpleVector(),
	}
}

func (g *Graph) Size() int {
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
