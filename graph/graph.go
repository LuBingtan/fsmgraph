package graph

import "fmt"

/*****************************************  graph interface  *****************************************/

// define for graph type
type GraphType string

// define for graph interface
type GraphInterface interface {
	/////// meta data ///////
	// update
	SetName(string)
	SetType(GraphType)
	// read
	Name() string
	Type() GraphType

	/////// relation data ///////
	// create
	InsertVertex(VertexInterface) (index int, err error)
	InsertEdge(src, dst VertexInterface, ei EdgeInterface) error
	// read
	GetVertex(id string) VertexInterface
	Verteces() map[string]VertexInterface
	// update
	SetVertex(v VertexInterface) error
	SetEdge(src, dst VertexInterface, ei EdgeInterface) error
	// delete
	RemoveVertex(VertexInterface)
	RemoveEdge(src, dst VertexInterface) error
}

/*****************************************  graph struct  *****************************************/
type Graph struct {
	name      string
	graphType GraphType
	verteces  map[string]VertexInterface
}

func NewGraph(name string) *Graph {
	return &Graph{
		name:     name,
		verteces: make(map[string]VertexInterface),
	}
}

// update graph name
func (g *Graph) SetName(n string) {
	g.name = n
}

// read graph name
func (g *Graph) Name() string {
	return g.name
}

// update graph type
func (g *Graph) SetType(t GraphType) {
	g.graphType = t
}

// read graph name
func (g *Graph) Type() GraphType {
	return g.graphType
}

// insert a new vertex
func (g *Graph) InsertVertex(v VertexInterface) error {
	if _, ok := g.verteces[v.Id()]; ok {
		return fmt.Errorf("vertex[id:%s] already exists!", v.Id())
	}

	g.verteces[v.Id()] = v

	return nil
}

// insert a new edge which is from src vertex to dst vertex
func (g *Graph) InsertEdge(src, dst VertexInterface, ei EdgeInterface) error {
	if _, ok := g.verteces[src.Id()]; !ok {
		return fmt.Errorf("vertex[id:%s] not exists, insert vertex first!", src.Id())
	}

	if _, ok := g.verteces[dst.Id()]; !ok {
		return fmt.Errorf("vertex[id:%s] not exists, insert vertex first!", dst.Id())
	}

	src.Adjoin(dst, ei)

	return nil
}

// get a vertex by id
func (g *Graph) GetVertex(id string) VertexInterface {
	v, ok := g.verteces[id]
	if !ok {
		return nil
	}

	return v
}

// get id-vertex map
func (g *Graph) Verteces() map[string]VertexInterface {
	return g.verteces
}

// update a vertex with specified id, and return an error if the id not exist
func (g *Graph) SetVertex(v VertexInterface) error {
	if _, ok := g.verteces[v.Id()]; !ok {
		return fmt.Errorf("vertex[id:%s] not exists, insert vertex first!", v.Id())
	}

	g.verteces[v.Id()] = v

	return nil
}

// update a edge which is from src verte to dst vertex
func (g *Graph) SetEdge(src, dst VertexInterface, ei EdgeInterface) error {
	if _, ok := g.verteces[src.Id()]; !ok {
		return fmt.Errorf("vertex[id:%s] not exists, insert vertex first!", src.Id())
	}

	if _, ok := g.verteces[dst.Id()]; !ok {
		return fmt.Errorf("vertex[id:%s] not exists, insert vertex first!", dst.Id())
	}

	src.SetEdge(dst, ei)

	return nil
}

// remove a vertex in graph, and those verteces which point to this vertex will also remove it
func (g *Graph) RemoveVertex(v VertexInterface) {
	for _, src := range g.verteces {
		v.RemoveAdjoin(src)
		src.RemoveAdjoin(v)
	}
	delete(g.verteces, v.Id())
}

// remove a edge in graph which is from src verte to dst vertex
func (g *Graph) RemoveEdge(src, dst VertexInterface) error {
	if _, ok := g.verteces[src.Id()]; !ok {
		return fmt.Errorf("vertex[id:%s] not exists, insert vertex first!", src.Id())
	}

	if _, ok := g.verteces[dst.Id()]; !ok {
		return fmt.Errorf("vertex[id:%s] not exists, insert vertex first!", dst.Id())
	}

	src.RemoveAdjoin(dst)

	return nil
}
