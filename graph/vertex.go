package graph

import (
	"fmt"
	simpleSt "fsmgraph-lib/simplestructure"
	"reflect"
)

/*****************************************  vertex interface  *****************************************/

// define for vertex type
type VertexType string

// define for vertex state stype
type VertexState string

// define for vertex executor function type
type ExecutorFunc func(...interface{}) (interface{}, error)

// define for vertex interface
type VertexInterface interface {
	/////// meta data ///////
	// update
	SetType(VertexType)
	SetId(string)
	SetData(interface{})
	SetState(VertexState)
	// read
	Type() VertexType
	Id() string
	Data() interface{}
	State() VertexState
	// behavior
	SetExecutor(ExecutorFunc)
	Execute(...interface{}) (interface{}, error)

	/////// relation data ///////
	// update
	Adjoin(dst VertexInterface, ei EdgeInterface)
	SetEdge(dst VertexInterface, ei EdgeInterface)
	incIndegree()
	decIndegree()
	incOutdegree()
	decOutdegree()
	// delete
	RemoveAdjoin(VertexInterface)
	// read
	FindAdjoinVertex(VertexInterface) int
	Edges() []EdgeInterface
	Indegree() int
	Outdegree() int
}

/*****************************************  vertex struct  *****************************************/

/// [Define]
// structure for vertex
type Vertex struct {
	// meta data
	vertexType VertexType
	id         string
	data       interface{}
	state      VertexState
	executor   ExecutorFunc
	// graph data
	edges     simpleSt.SimpleVector
	indegree  int
	outdegree int
}

func NewVertex(id string, data interface{}) *Vertex {
	return &Vertex{
		id:   id,
		data: data,
	}
}

// Update vertex type
func (v *Vertex) SetType(t VertexType) {
	v.vertexType = t
}

// Update vertex id
func (v *Vertex) SetId(id string) {
	v.id = id
}

// Update vertex data
func (v *Vertex) SetData(data interface{}) {
	v.data = data
}

// Update vertex state
func (v *Vertex) SetState(state VertexState) {
	v.state = state
}

// get vertex type
func (v *Vertex) Type() VertexType {
	return v.vertexType
}

// get vertex id
func (v *Vertex) Id() string {
	return v.id
}

// get vertex data
func (v *Vertex) Data() interface{} {
	return v.data
}

// get vertex state
func (v *Vertex) State() VertexState {
	return v.state
}

// vertex behavior set
func (v *Vertex) SetExecutor(executorFunc ExecutorFunc) {
	v.executor = executorFunc
}

// vertex behavior execute
func (v *Vertex) Execute(inputs ...interface{}) (interface{}, error) {
	if v.executor == nil {
		return nil, fmt.Errorf("no executor.")
	}

	return v.executor(inputs)
}

// update adjacent vertex
func (v *Vertex) Adjoin(dst VertexInterface, ei EdgeInterface) {
	ei.SetVertex(dst)
	v.edges.Pushback(dst)
	v.incOutdegree()
	dst.incIndegree()
}

// update edge
func (v *Vertex) SetEdge(dst VertexInterface, ei EdgeInterface) {
	v.RemoveAdjoin(dst)
	v.Adjoin(dst, ei)
}

// increase indegree
func (v *Vertex) incIndegree() {
	v.indegree++
}

// decrease indegree
func (v *Vertex) decIndegree() {
	v.indegree--
}

// increase outdegree
func (v *Vertex) incOutdegree() {
	v.outdegree++
}

// decrease outdegree
func (v *Vertex) decOutdegree() {
	v.outdegree--
}

// delete adjacent vertex
func (v *Vertex) RemoveAdjoin(vi VertexInterface) {
	index := v.FindAdjoinVertex(vi)
	if index == -1 {
		return
	}

	v.edges.Remove(index)
	v.decOutdegree()
	vi.decIndegree()
}

// find adjacent vertex and return its binding edge's index
func (v *Vertex) FindAdjoinVertex(vi VertexInterface) int {
	for i, d := range v.edges.Data() {
		e := d.(EdgeInterface)
		if reflect.DeepEqual(e.Vertex(), vi) {
			return i
		}
	}

	return -1
}

// get all edges
func (v *Vertex) Edges() (ei []EdgeInterface) {
	for _, d := range v.edges.Data() {
		ei = append(ei, d.(EdgeInterface))
	}

	return ei
}

// Indegree
func (v *Vertex) Indegree() int {
	return v.indegree
}

// outdegree
func (v *Vertex) Outdegree() int {
	return v.outdegree
}
