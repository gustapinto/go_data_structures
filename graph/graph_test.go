package graph

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	expected := &Graph{
		Nodes: []*GraphNode{
			{Value: 10},
		},
	}

	graph := NewGraph(&GraphNode{10})

	if !reflect.DeepEqual(graph, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, graph)
	}
}

func TestSizeMustBeZeroOnNil(t *testing.T) {
	size := Size(nil)

	if size != 0 {
		t.Errorf("Failed! Expected size to be 0, got %d", size)
	}
}

func TestSizeMustReturnTheActualSize(t *testing.T) {
	graph := NewGraph(&GraphNode{10})
	AddNode(graph, &GraphNode{4})
	AddNode(graph, &GraphNode{6})

	size := Size(graph)

	if size != 3 {
		t.Errorf("Failed! Expected size to be 3, got %d", size)
	}
}

func TestAddNodeMustAddANode(t *testing.T) {
	expected := &Graph{
		Nodes: []*GraphNode{
			{Value: 10},
			{Value: 2},
		},
	}

	graph := NewGraph(&GraphNode{10})
	AddNode(graph, &GraphNode{2})

	if !reflect.DeepEqual(graph, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, graph)
	}
}

func TestAddEdge(t *testing.T) {
	node1 := &GraphNode{10}
	node2 := &GraphNode{2}
	node3 := &GraphNode{4}

	expected := &Graph{
		Nodes: []*GraphNode{node1, node2, node3},
		Edges: map[GraphNode][]*GraphNode{
			*node1: {node2, node3},
			*node2: {node1},
			*node3: {node1},
		},
	}

	graph := NewGraph(node1)
	AddNode(graph, node2)
	AddNode(graph, node3)
	AddEdge(graph, node1, node2)
	AddEdge(graph, node1, node3)

	if !reflect.DeepEqual(graph, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, graph)
	}
}
