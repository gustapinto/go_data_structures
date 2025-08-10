package orderedmap

import (
	"slices"
	"testing"
)

func makeTestMap() *Map[string, string] {
	orderedMap := NewMap[string, string]()
	orderedMap.Set("key1", "value1")
	orderedMap.Set("key2", "value2")
	orderedMap.Set("key3", "value3")

	return orderedMap
}

func TestSetMustAddAllKeyValuePairs(t *testing.T) {
	orderedMap := NewMap[string, string]()
	orderedMap.Set("key1", "value1")
	orderedMap.Set("key2", "value2")
	orderedMap.Set("key3", "value3")
	orderedMap.Set("key4", "value4")
	orderedMap.Set("key5", "value5")

	expectedLen := uint(5)
	actualLen := orderedMap.Len()

	if expectedLen != actualLen {
		t.Errorf("expected [%d], got [%d]", expectedLen, actualLen)
	}
}

func TestIterMustIterateInOrder(t *testing.T) {
	orderedMap := makeTestMap()

	expected := []string{"key1", "key2", "key3"}
	actual := []string{}

	orderedMap.Iter(func(key, _ string) bool {
		actual = append(actual, key)

		return false
	})

	if slices.Equal(expected, actual) {
		t.Errorf("expected [%v], got [%v]", expected, actual)
	}
}

func TestIterMustStopEarly(t *testing.T) {
	orderedMap := makeTestMap()

	expected := []string{"key1"}
	actual := []string{}

	orderedMap.Iter(func(key, _ string) bool {
		actual = append(actual, key)

		return true
	})

	if len(expected) != len(actual) {
		if slices.Equal(expected, actual) {
			t.Errorf("expected [%v], got [%v]", expected, actual)
		}
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("expected [%v], got [%v]", expected, actual)
		}
	}
}

func TestGetMustFindTheKey(t *testing.T) {
	orderedMap := makeTestMap()

	expected := "value2"
	actual, _ := orderedMap.Get("key2")

	if expected != actual {
		t.Errorf("expected [%v], got [%v]", expected, actual)
	}
}
