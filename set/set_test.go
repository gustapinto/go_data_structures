package set

import (
	"reflect"
	"testing"
)

func TestCreateMustCreateANewSet(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{10: true}}

	set := Create(10)

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, set)
	}
}

func TestContainsMustNegateOnNil(t *testing.T) {
	contains := Contains[int](nil, 10)

	if contains {
		t.Errorf("Failed! Expected to fail, got %v", contains)
	}
}

func TestConatins(t *testing.T) {
	set := Create(10)
	Push(set, 4)
	Push(set, 6)

	contains := Contains(set, 6)

	if !contains {
		t.Errorf("Failed! Expected true, got %v", contains)
	}
}

func TestSizeMustBeZeroOnNil(t *testing.T) {
	size := Size[int](nil)
	if size != 0 {
		t.Errorf("Failed! Expected size to be zero, got %d", size)
	}
}

func TestSize(t *testing.T) {
	set := Create(10)
	Push(set, 4)
	Push(set, 6)
	Push(set, -2)
	Push(set, 2)

	size := Size(set)
	if size != 5 {
		t.Errorf("Failed! Expected size to be 5, got %d", size)
	}
}

func TestPushMustAddAItemToTheSet(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{10: true, 4: true, 6: true}}

	set := Create(10)
	Push(set, 4)
	Push(set, 6)

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, set)
	}
}

func TestPushMustNotAddDuplicatedItens(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{10: true, 4: true}}

	set := Create(10)
	Push(set, 4)
	Push(set, 4)
	Push(set, 4)

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, set)
	}
}

func TestPushMustHandleNilSetValues(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{10: true}}
	set := &Set[int]{}
	Push(set, 10)

	if !reflect.DeepEqual(set, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, set)
	}
}

func TestTraverseMustReturnEmptyOnNil(t *testing.T) {
	expected := []int{}
	values := Traverse[int](nil)

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}

func TestTraverMustReturnTheSetInnerValues(t *testing.T) {
	expected := []int{10, 4, 6}

	set := Create(10)
	Push(set, 4)
	Push(set, 6)

	values := Traverse[int](set)

	if len(values) != 3 || !Contains(set, 10) || !Contains(set, 4) || !Contains(set, 6) {
		t.Errorf("Failed! Expected %v, got %v", expected, values)
	}
}

func TestUnionMustReturnNilWhenBotAreNil(t *testing.T) {
	union := Union[int](nil, nil)

	if union != nil {
		t.Errorf("Failed! Expected nil, got %v", union)
	}
}

func TestUnionMustReturnThePAssedSetWhenOnlyOneIsNil(t *testing.T) {
	set := Create(10)
	Push(set, 4)

	union := Union(set, nil)
	if !reflect.DeepEqual(set, union) {
		t.Errorf("Failed! Expected %v, got %v", set, union)
	}

	union2 := Union(nil, set)
	if !reflect.DeepEqual(set, union2) {
		t.Errorf("Failed! Expected %v, got %v", set, union2)
	}
}

func TestUnionMustReturnTheUnionBetweenSetsWithoutDuplicates(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{10: true, 4: true, 6: true, -2: true}}

	set1 := Create(10)
	Push(set1, 4)
	Push(set1, 6)

	set2 := Create(-2)
	Push(set2, 4)

	union := Union(set1, set2)

	if !reflect.DeepEqual(union, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, union)
	}
}

func TestIntersectionMustReturnNilWhenAnySetIsNil(t *testing.T) {
	set := Create(10)
	Push(set, 4)

	intersection := Intersection(set, nil)
	if intersection != nil {
		t.Errorf("Failed! Expected nil, got %v", intersection)
	}

	intersection2 := Intersection(nil, set)
	if intersection2 != nil {
		t.Errorf("Failed! Expected nil, got %v", intersection)
	}
}

func TestIntersectionMustReturnTheIntersectionBetweenTwoNonNilSets(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{10: true, 4: true}}

	set1 := Create(10)
	Push(set1, 4)
	Push(set1, 6)

	set2 := Create(10)
	Push(set2, -2)
	Push(set2, 4)

	intersection := Intersection(set1, set2)

	if !reflect.DeepEqual(intersection, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, intersection)
	}
}

func TestIntersectionMustIndependOfSetOrder(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{10: true, 4: true}}

	set1 := Create(10)
	Push(set1, 4)
	Push(set1, 6)

	set2 := Create(10)
	Push(set2, -2)
	Push(set2, 4)

	intersection := Intersection(set2, set1)

	if !reflect.DeepEqual(intersection, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, intersection)
	}
}

func TestDifferenceMustReturnNilWhenTheFirstSetISNil(t *testing.T) {
	set := Create(10)

	difference := Difference(nil, set)
	if difference != nil {
		t.Errorf("Failed! Expected nil, got %v", difference)
	}
}

func TestDifferenceMustReturnTheFirstSetWhenTheSecondIsNil(t *testing.T) {
	set := Create(10)

	difference := Difference(set, nil)
	if !reflect.DeepEqual(set, difference) {
		t.Errorf("Failed! Expected %v, got %v", difference, set)
	}
}

func TestDifferenceMustReturnTheElementsFromTheFirstSetThatDontExistInTheSecond(t *testing.T) {
	expected := &Set[int]{Values: map[int]bool{6: true}}

	set1 := Create(10)
	Push(set1, 4)
	Push(set1, 6)

	set2 := Create(10)
	Push(set2, 4)

	difference := Difference(set1, set2)
	if !reflect.DeepEqual(difference, expected) {
		t.Errorf("Failed! Expected %v, got %v", expected, difference)
	}
}

func TestIsSubsetMustReturnFalseOnAnyNil(t *testing.T) {
	set := Create(10)

	isSubset := IsSubset(nil, set)
	if isSubset {
		t.Errorf("Failed! Expected false, got %v", isSubset)
	}

	isSubset2 := IsSubset(set, nil)
	if isSubset2 {
		t.Errorf("Failed! Expected false, got %v", isSubset2)
	}
}

func TestIsSusetMustReturnFalseWhenSet1IsGreaterThanSet2(t *testing.T) {
	set1 := Create(10)
	Push(set1, 4)

	set2 := Create(10)

	isSubset := IsSubset(set1, set2)
	if isSubset {
		t.Errorf("Failed! Expected %v to not be a subset of %v", set1, set2)
	}
}

func TestIsSubsetMustReturnTrueWhenSet1IsSubsetOfSet2(t *testing.T) {
	set1 := Create(10)
	Push(set1, 4)

	set2 := Create(10)
	Push(set2, 4)
	Push(set2, 6)
	Push(set2, -2)

	isSubset := IsSubset(set1, set2)
	if !isSubset {
		t.Errorf("Failed! Expected %v to be subset of %v", set1, set2)
	}
}

func TestIsDisjointMustReturnFalseOnAnyNil(t *testing.T) {
	isDisjoint := IsDisjoint[int](nil, nil)
	if isDisjoint {
		t.Errorf("Failed! Expected to be false on nil, got %v", isDisjoint)
	}
}

func TestIsDisjointMustReturnFalseIfAnyItemIsPresentOnBothSets(t *testing.T) {
	set1 := Create(10)
	Push(set1, 4)
	Push(set1, 6)
	Push(set1, -2)

	set2 := Create(4)

	isDisjoint := IsDisjoint(set1, set2)
	if isDisjoint {
		t.Errorf("Failed! Expected %v and %v to do not be disjoint", set1, set2)
	}
}

func TestIsDisjointMustReturnTrueWhenBothSetsDoesNotShareAnyItems(t *testing.T) {
	set1 := Create(10)
	Push(set1, 4)
	Push(set1, 6)

	set2 := Create(20)
	Push(set2, 8)
	Push(set2, 12)
	Push(set2, -4)

	isDisjoint := IsDisjoint(set1, set2)
	if !isDisjoint {
		t.Errorf("Failed! Expected %v and %v to be disjoint", set1, set2)
	}
}
