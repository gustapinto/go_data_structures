package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestAddNodeMustCreateANewDoublyLinkedList(t *testing.T) {
	expected := &ListNode[int]{Value: 10}

	node := AddNode(nil, 10)

	if !reflect.DeepEqual(node, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, node)
	}
}

func TestAddNodeMustAddANewNode(t *testing.T) {
	expected := &ListNode[int]{
		Value: 10,
	}
	expected.Next = &ListNode[int]{
		Value:    5,
		Previous: expected,
	}

	node := AddNode(nil, 10)
	AddNode(node, 5)

	if !reflect.DeepEqual(node, expected) {
		t.Errorf("Failed! Expected %#v, got %#v", expected, node)
	}
}

func TestTraverseMustReturnAEmptySliceOnEmptyNode(t *testing.T) {
	expected := []int{}
	values := Traverse[int](nil)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %#v, got %#v", expected, values)
	}
}

func TestTraverseMustReturnAllValues(t *testing.T) {
	expected := []int{1, -2, -3, 4, 5}

	node := AddNode(nil, 1)
	AddNode(node, -2)
	AddNode(node, -3)
	AddNode(node, 4)
	AddNode(node, 5)

	values := Traverse(node)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %#v, got %#v", expected, values)
	}
}

func TestReverseMustReturnAllValuesInReversedOrder(t *testing.T) {
	expected := []int{5, 4, -3, -2, 1}

	node := AddNode(nil, 1)
	AddNode(node, -2)
	AddNode(node, -3)
	AddNode(node, 4)
	AddNode(node, 5)

	values := Reverse(node)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %#v, got %#v", expected, values)
	}
}

func TestSize(t *testing.T) {
	node := AddNode(nil, 1)
	AddNode(node, -2)
	AddNode(node, -3)
	AddNode(node, 4)
	AddNode(node, 5)

	size := Size(node)

	if size != 5 {
		t.Errorf("Failed! Expected size 5, got %d", size)
	}
}

func TestLookupOnNil(t *testing.T) {
	exists := Lookup(nil, 10)

	if exists {
		t.Errorf("Failed! Expected to not find value on nil")
	}
}

func TestLookup(t *testing.T) {
	node := AddNode(nil, 10)
	exists := Lookup(node, 10)

	if !exists {
		t.Errorf("Failed! Expected to find existing value")
	}
}
