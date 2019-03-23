package graph

/*****************************************  edge interface  *****************************************/

// define for edge type
type EdgeType string

// define for edge state
type EdgeState string

// define for edge interface
type EdgeInterface interface {
	/////// meta data ///////
	// update
	SetType(EdgeType)
	SetWeight(int)
	SetState(EdgeState)
	// read
	Type() EdgeType
	Weight() int
	State() EdgeState

	/////// relation data ///////
	// update
	SetVertex(VertexInterface)
	// read
	Vertex() VertexInterface
}

/*****************************************  edge struct  *****************************************/

type Edge struct {
	// meta data
	edgeType EdgeType
	weight   int
	state    EdgeState
	// graph data
	vertex VertexInterface
}

func NewEdge() *Edge {
	return &Edge{}
}

func (e *Edge) SetType(t EdgeType) {
	e.edgeType = t
}

func (e *Edge) SetWeight(w int) {
	e.weight = w
}

func (e *Edge) SetState(s EdgeState) {
	e.state = s
}

func (e *Edge) Type() EdgeType {
	return e.edgeType
}

func (e *Edge) Weight() int {
	return e.weight
}

func (e *Edge) State() EdgeState {
	return e.state
}

func (e *Edge) SetVertex(v VertexInterface) {
	e.vertex = v
}

func (e *Edge) Vertex() VertexInterface {
	return e.vertex
}
