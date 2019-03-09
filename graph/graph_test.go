package graph

import (
	"testing"
)

func Test4Graph4Init(t *testing.T) {
	t.Logf("testing for graph initialize start.")
	g := Graph{}
	v := Vertex{}
	vList := make([]interface{}, 0)
	vList = append(vList, v)
	g.Initialize(vList)
}

func Test4Graph4Traverse(t *testing.T) {
	t.Logf("testing for graph tranverse start.")
}
