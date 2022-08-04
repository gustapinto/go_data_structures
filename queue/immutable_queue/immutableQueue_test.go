package immutablequeue

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewImmutableQueue(t *testing.T) {
	expected := ImmutableQueue[int]{Values: []int{10, 4, 6, -2}}
	queue := NewImmutableQueue(10, 4, 6, -2)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}
}

func TestSize(t *testing.T) {
	queue := NewImmutableQueue(10, 4, 6, -2)
	size := Size(queue)

	if size != 4 {
		t.Errorf("Failed! Expected size to be 4, got %d", size)
	}
}

func TestPushMustAddANewItemToQueue(t *testing.T) {
	expected := ImmutableQueue[int]{Values: []int{10, 4, 6, -2, 2}}
	queue := NewImmutableQueue(10, 4)

	newQueue := Push(queue, 6, -2, 2)

	if !reflect.DeepEqual(newQueue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, newQueue)
	}
}

func TestPushMustNotMutateTheOriginalQueue(t *testing.T) {
	expected := ImmutableQueue[int]{Values: []int{10, 4}}
	queue := NewImmutableQueue(10, 4)

	Push(queue, 6, -2, 2)

	if !reflect.DeepEqual(queue, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, queue)
	}
}

func TestPopMustReturnErrTheZeroValues(t *testing.T) {
	queue := ImmutableQueue[int]{Values: []int{}}

	_, _, err := Pop(queue)
	if !errors.Is(err, ErrCannotPopOnEmptyQueue) {
		t.Errorf("Failed! Expected pop to error on empty queue")
	}
}

func TestPopMustReturnAEmptyQueueWhenItRemoveTheLastElement(t *testing.T) {
	expectedValue := 10
	expectedNewQueue := ImmutableQueue[int]{Values: []int{}}

	queue := NewImmutableQueue(10)
	value, newQueue, err := Pop(queue)

	if err != nil {
		t.Errorf("Failed! Did not expect to fail, got err %v", err)
	}
	if value != expectedValue {
		t.Errorf("Failed! Expected %d, got %d", expectedValue, value)
	}
	if !reflect.DeepEqual(newQueue, expectedNewQueue) {
		t.Errorf("Failed! Expected %v, got %v", expectedNewQueue, newQueue)
	}
}

func TestPopMustReturnTheFirstElementAndANewQueue(t *testing.T) {
	expectedValue := 10
	expectedNewQueue := ImmutableQueue[int]{Values: []int{4, 6}}

	queue := NewImmutableQueue(10, 4, 6)
	value, newQueue, err := Pop(queue)

	if err != nil {
		t.Errorf("Failed! Did not expect to fail, got err %v", err)
	}
	if value != expectedValue {
		t.Errorf("Failed! Expected %d, got %d", expectedValue, value)
	}
	if !reflect.DeepEqual(newQueue, expectedNewQueue) {
		t.Errorf("Failed! Expected %v, got %v", expectedNewQueue, newQueue)
	}
}
