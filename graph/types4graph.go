package graph

import (
	"reflect"
)

type GraphType string

type GraphInterface interface {
	Initialize([]interface{}) error
	Tranverse() []*Vertex
	Type() GraphType
}

type ExecutorFunc func(interface{}) (interface{}, error)

type VertexState string

/// [Define]
// structure for vertex
type Vertex struct {
	Id        int
	Data      interface{}
	FirstEdge *Edge
	State     VertexState
	Executor  ExecutorFunc
}

type Edge struct {
	VertexPoint *Vertex
	NextEdge    *Edge
}

type Graph struct {
	Vertexlist []*Vertex
}

func (g *Graph) Initialize(vertexes []interface{}) error {
	for _, v := range vertexes {
		// type check
		if tp := reflect.TypeOf(v); tp.Name() != "Vertex" {
			return TypeError
		}

	}
	return nil
}
