package graph



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
	// update
	InsertVertex(VertexInterface) (index int, err error)
	InsertEdge(a, b VertexInterface) error
	// read
	GetVertexById(id string)
	GetVertexByIndex(index int)
	Verteces() []VertexInterface
	Edges() []EdgeInterface
	// update
	SetVertexById(id string, v VertexInterface)
	SetVertexByIndex(index int, v VertexInterface)
	// delete
	RemoveVertex(VertexInterface) error

	/////// algorithms ///////
	TopoSort() ([]VertexInterface, error)
	BFS() []VertexInterface
	DFS() []VertexInterface
}

/*****************************************  graph struct  *****************************************/