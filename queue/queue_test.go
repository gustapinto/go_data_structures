package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	expected := &QueueNode{
		Value: 10,
		Next: nil,
	}

	queue := NewQueue(10)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}
}

func TestSizeMustReturnZeroOnNil(t *testing.T) {
	size := Size(nil)

	if size != 0 {
		t.Errorf("Failed! Expected size to be 0, got %d", size)
	}
}

func TestSizeMustReturnTheSizeOfQueue(t *testing.T) {
	queue := &QueueNode{
		Value: 10,
		Next: &QueueNode{
			Value: 4,
			Next: &QueueNode{
				Value: -2,
				Next:  nil,
			},
		},
	}

	size := Size(queue)

	if size != 3 {
		t.Errorf("Failed! Expected size to be 3, got %d", size)
	}
}

func TestSizeMustNotMutateTheQueue(t *testing.T) {
	queue := &QueueNode{
		Value: 10,
		Next: &QueueNode{
			Value: 4,
			Next: &QueueNode{
				Value: -2,
				Next:  nil,
			},
		},
	}

	original := queue

	_ = Size(queue)

	if !reflect.DeepEqual(queue, original) {
		t.Errorf("Failed! Expected size to do not mutate the queue, got %v", queue)
	}
}

func TestPushMustInitializeANewQueue(t *testing.T) {
	expected := &QueueNode{
		Value: 10,
		Next:  nil,
	}

	queue := Push(nil, 10)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}
}

func TestPushMustAddANewItem(t *testing.T) {
	expected := &QueueNode{
		Value: 10,
		Next: &QueueNode{
			Value: 4,
			Next: &QueueNode{
				Value: -2,
				Next:  nil,
			},
		},
	}

	queue := Push(nil, 10)
	Push(queue, 4)
	Push(queue, -2)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}
}

func TestPopMustReturnNilOnNil(t *testing.T) {
	value := Pop(nil)

	if value != nil {
		t.Errorf("Failed! Expected nil, got %v", value)
	}
}

func TestPopMustDefineQueueAsItZeroValueWhenItOnlyHaveOneItem(t *testing.T) {
	expected := &QueueNode{}

	queue := &QueueNode{
		Value: 4,
		Next:  nil,
	}

	_ = Pop(queue)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}
}

func TestPopMustReturnTheFirstValue(t *testing.T) {
	queue := &QueueNode{
		Value: 10,
		Next: &QueueNode{
			Value: 4,
			Next:  nil,
		},
	}

	value := Pop(queue)

	if value != 10 {
		t.Errorf("Failed! Expected value to be 10, got %d", value)
	}
}

func TestPopMustRemoveAItem(t *testing.T) {
	expected := &QueueNode{
		Value: 4,
		Next:  nil,
	}

	queue := &QueueNode{
		Value: 10,
		Next: &QueueNode{
			Value: 4,
			Next:  nil,
		},
	}

	_ = Pop(queue)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}

	size := Size(queue)

	if size != 1 {
		t.Errorf("Failed! Expected size to be 1, got %d", size)
	}
}

func TestTraverseMustReturnAEmptySliceOnNi(t *testing.T) {
	values := Traverse(nil)

	if !reflect.DeepEqual(values, []any{}) {
		t.Errorf("Failed! Expected a empty slice, got %v", values)
	}
}

func TestTraverseMustRrturnTheQueueValues(t *testing.T) {
	expected := []any{10, 4, 6}

	queue := Push(nil, 10)
	Push(queue, 4)
	Push(queue, 6)

	values := Traverse(queue)

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}
