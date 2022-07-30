package linkedlist

import (
	"reflect"
	"testing"
)

func TestAddNodeMustCreateANewNode(t *testing.T) {
	expected := &ListNode[int]{Value: 10}

	node := AddNode(nil, 10)

	if !reflect.DeepEqual(node, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, node)
	}
}

func TestAddNodeWithSameValueMustReturnTheSameNode(t *testing.T) {
	node := &ListNode[int]{Value: 10}
	newNode := AddNode(node, 10)

	if node != newNode {
		t.Errorf("Failed! Expected %v, got %v", node, newNode)
	}
}

func TestAddNodeMustAddANewNode(t *testing.T) {
	expected := &ListNode[int]{
		Value: 10,
		Next: &ListNode[int]{
			Value: 5,
		},
	}

	node := AddNode(nil, 10)
	AddNode(node, 5)

	if !reflect.DeepEqual(node, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, node)
	}
}

func TestTraverseMustReturnAEmptySliceOnEmptyNode(t *testing.T) {
	expected := []int{}
	values := Traverse[int](nil)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %#v, got %#v", expected, values)
	}
}

func TestTraveseShouldReturnASliceOfValeus(t *testing.T) {
	expected := []int{10, 2, 4, 1, -6}

	node := AddNode(nil, 10)
	AddNode(node, 2)
	AddNode(node, 4)
	AddNode(node, 1)
	AddNode(node, -6)

	values := Traverse(node)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}

func TestLookupMustFailOnNilNode(t *testing.T) {
	exists := Lookup(nil, 10)

	if exists {
		t.Errorf("Faield! Expected to fail on nil node")
	}
}

func TestLookupMustReturnTrueWhenValueExists(t *testing.T) {
	node := AddNode(nil, 10)
	exists := Lookup(node, 10)

	if !exists {
		t.Errorf("Faield! Expected to fail on nil node")
	}
}

func TestSize(t *testing.T) {
	node := AddNode(nil, 10)
	AddNode(node, 2)
	AddNode(node, 4)
	AddNode(node, 1)
	AddNode(node, -6)

	size := Size(node)

	if size != 5 {
		t.Errorf("Failed! Expected 5, got %d", size)
	}
}
