package stack

import (
	"reflect"
	"testing"
)

func TestSizeMustBeZeroOnNil(t *testing.T) {
	size := Size(nil)

	if size != 0 {
		t.Errorf("Failed! Expected 0, got %d", size)
	}
}

func TestSizeMusCountAllnodes(t *testing.T) {
	stack := &StackNode{
		Value: 10,
		Next: &StackNode{
			Value: 4,
			Next: &StackNode{
				Value: 6,
				Next:  nil,
			},
		},
	}

	size := Size(stack)

	if size != 3 {
		t.Errorf("Failed! Expected 3, got %d", size)
	}
}

func TestPushMustCreateANewStack(t *testing.T) {
	expected := &StackNode{Value: 10}

	stack := Push(nil, 10)

	if !reflect.DeepEqual(stack, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, stack)
	}
}

func TestPushMustAddANewItemOnStackStart(t *testing.T) {
	expected := &StackNode{
		Value: 10,
		Next: &StackNode{
			Value: 4,
			Next: &StackNode{
				Value: 6,
				Next:  nil,
			},
		},
	}

	stack := Push(nil, 6)
	Push(stack, 4)
	Push(stack, 10)

	if !reflect.DeepEqual(stack, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, stack.Next)
	}
}

func TestPopMustReturnNilOnNil(t *testing.T) {
	item := Pop(nil)

	if item != nil {
		t.Errorf("Failed! Expected nil, got %v", item)
	}
}

func TestPopMustReturnTheLastValue(t *testing.T) {
	stack := &StackNode{
		Value: 10,
		Next: &StackNode{
			Value: 4,
			Next: &StackNode{
				Value: 6,
				Next:  nil,
			},
		},
	}

	item := Pop(stack)

	if item != 10 {
		t.Errorf("Failed! Expected value to be 10, got %v", item)
	}
}

func TestPopMustRemoveAItem(t *testing.T) {
	expected := &StackNode{
		Value: 4,
		Next:  nil,
	}
	stack := &StackNode{
		Value: 10,
		Next: &StackNode{
			Value: 4,
			Next:  nil,
		},
	}

	_ = Pop(stack)

	if !reflect.DeepEqual(stack, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, stack)
	}

	size := Size(stack)

	if size != 1 {
		t.Errorf("Failed! Expected size to be 1, got %d", size)
	}
}
