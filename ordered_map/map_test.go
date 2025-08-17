package orderedmap

import (
	"fmt"
	"slices"
	"testing"
)

func makeTestMap(size uint) *Map[string, string] {
	orderedMap := NewMap[string, string]()

	for i := range size {
		key := fmt.Sprintf("key%d", i+1)
		value := fmt.Sprintf("value%d", i+1)

		orderedMap.Set(key, value)
	}

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
	orderedMap := makeTestMap(5)

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
	orderedMap := makeTestMap(5)

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

func TestGetMustFindAExistingKey(t *testing.T) {
	orderedMap := makeTestMap(5)

	expected := "value2"
	expectedExists := true
	actual, actualExists := orderedMap.Get("key2")

	if expected != actual || expectedExists != actualExists {
		t.Errorf("expected [%v], got [%v]", expected, actual)
	}
}

func TestGetMustNotFindAExistingKey(t *testing.T) {
	orderedMap := NewMap[string, string]()

	if _, exists := orderedMap.Get("invalid_key"); exists {
		t.Error("expected [invalid_key] to no exists")
	}
}

func TestDelMustRemoveTheKeyFromTheMap(t *testing.T) {
	orderedmap := makeTestMap(5)
	orderedmap.Del("key1")

	if orderedmap.Len() != 4 {
		t.Error("expected len to be 4")
	}

	orderedmap.Del("key4")
	if orderedmap.Len() != 3 {
		t.Error("expected len to be 3")
	}

	if _, exists := orderedmap.Get("key2"); !exists {
		t.Error("expected [key2] to exists")
	}
}
