package slicequeue

import (
	"reflect"
	"testing"
)

func TestNewSliceQueue(t *testing.T) {
	expected := &SliceQueue{
		Values: []any{10},
	}

	queue := NewSliceQueue(10)

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

func SizeMustReturnTheActualSize(t *testing.T) {
	queue := Push(nil, 10)
	Push(queue, 4)
	Push(queue, 6)

	size := Size(queue)

	if size != 3 {
		t.Errorf("Failed! Expected size to be 3, got %d", size)
	}
}

func TestPushMustCreateANewQueueOnNil(t *testing.T) {
	expected := &SliceQueue{
		Values: []any{10},
	}

	queue := Push(nil, 10)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}
}

func TestPushMustAddAItemToTheQueueValues(t *testing.T) {
	expected := []any{10, 4, 6, 2, -2}

	queue := Push(nil, 10)
	Push(queue, 4)
	Push(queue, 6)
	Push(queue, 2)
	Push(queue, -2)

	if !reflect.DeepEqual(queue.Values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue.Values)
	}
}

func TestPopMustReturnEmptyOnNil(t *testing.T) {
	expected := []any{}
	value := Pop(nil)

	if !reflect.DeepEqual(value, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, value)
	}
}

func TestPopMustMustReturnTheFirstValueOfQueue(t *testing.T) {
	queue := Push(nil, 10)
	Push(queue, 4)
	Push(queue, 6)

	value := Pop(queue)

	if value != 10 {
		t.Errorf("Failed! Expected value to be 10, got %v", value)
	}
}

func TestPopMustRemoveAElementFromQueueValues(t *testing.T) {
	expected := []any{4, 6}

	queue := Push(nil, 10)
	Push(queue, 4)
	Push(queue, 6)

	_ = Pop(queue)

	if !reflect.DeepEqual(queue.Values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue.Values)
	}
}

func TestTraverseMustReturnEmptyOnNil(t *testing.T) {
	expected := []any{}

	values := Traverse(nil)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}

func TestTraverseMustReturnTheInnerSliceOfQueue(t *testing.T) {
	queue := Push(nil, 10)
	Push(queue, 4)
	Push(queue, 6)

	values := Traverse(queue)

	if !reflect.DeepEqual(values, queue.Values) {
		t.Errorf("Failed! Expected %v, got %v", queue.Values, values)
	}
}
