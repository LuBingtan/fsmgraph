package fsm

import (
	"fsmgraph-lib/graph"
)

// struct FSM
// A FSM is a graph, its' vertex is state. But a FSM is also a vertex for a 'FSMGraph'.
type FSM struct {
	graph.AbstractGraph
	graph.AbstractVertex
}
