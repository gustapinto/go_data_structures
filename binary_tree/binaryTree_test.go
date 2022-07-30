package binary_tree

import (
	"reflect"
	"testing"
)

func TestInsertMustCreateANewTree(t *testing.T) {
	expected := &BinaryTree[int]{nil, 4, nil}

	tree := Insert(nil, 4)

	if !reflect.DeepEqual(expected, tree) {
		t.Errorf("Failed! Expected %v, got %v", expected, tree)
	}
}

func TestInsertMustOrderInsertions(t *testing.T) {
	expected := &BinaryTree[int]{&BinaryTree[int]{nil, 2, nil}, 4, &BinaryTree[int]{nil, 10, nil}}

	tree := Insert(nil, 4)
	Insert(tree, 10)
	Insert(tree, 2)

	if !reflect.DeepEqual(tree, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, tree)
	}
}

func TestTraverse(t *testing.T) {
	expected := []int{-10, -2, 2, 4, 6}

	tree := Insert(nil, 4)
	Insert(tree, -10)
	Insert(tree, 6)
	Insert(tree, 2)
	Insert(tree, -2)

	values := Traverse(tree)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}

func TestTraversePostOrder(t *testing.T) {
	expected := []int{-2, 2, -10, 6, 4}

	tree := Insert(nil, 4)
	Insert(tree, -10)
	Insert(tree, 6)
	Insert(tree, 2)
	Insert(tree, -2)

	values := TraversePostOrder(tree)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}

func TestTraversePreOrder(t *testing.T) {
	expected := []int{4, -10, 2, -2, 6}

	tree := Insert(nil, 4)
	Insert(tree, -10)
	Insert(tree, 6)
	Insert(tree, 2)
	Insert(tree, -2)

	values := TraversePreOrder(tree)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}
