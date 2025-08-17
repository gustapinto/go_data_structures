package orderedmap

import (
	"sync"
)

// Map impl
type Map[K comparable, V any] struct {
	size    uint
	entries *Entry[K, V]
	mu      sync.Mutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		size:    0,
		entries: nil,
		mu:      sync.Mutex{},
	}
}

func (om *Map[K, V]) Len() uint {
	return om.size
}

func (om *Map[K, V]) Set(key K, value V) {
	om.mu.Lock()
	defer om.mu.Unlock()

	entry := &Entry[K, V]{
		Key:   key,
		Value: value,
	}

	if om.entries == nil {
		om.entries = entry
		om.size++
		return
	}

	om.entries.append(entry)
	om.size++
}

func (om *Map[K, V]) Get(key K) (value V, exists bool) {
	if om.entries == nil || om.size == 0 {
		return value, false
	}

	om.mu.Lock()
	defer om.mu.Unlock()

	om.entries.iter(func(me *Entry[K, V]) (stop bool) {
		if me.Key == key {
			value = me.Value
			exists = true
			return true
		}

		return false
	})

	return value, exists
}

func (om *Map[K, V]) Exists(key K) bool {
	_, exists := om.Get(key)

	return exists
}

func (om *Map[K, V]) Iter(fn func(key K, value V) (stop bool)) {
	if om.entries == nil {
		return
	}

	om.mu.Lock()
	defer om.mu.Unlock()

	om.entries.iter(func(me *Entry[K, V]) bool {
		return fn(me.Key, me.Value)
	})
}

func (om *Map[K, V]) Del(key K) {
	if om.entries == nil || om.size == 0 {
		return
	}

	om.mu.Lock()
	defer om.mu.Unlock()

	om.entries.remove(key)
	om.size--
}

// Entry impl
type Entry[K comparable, V any] struct {
	Key   K
	Value V
	prev  *Entry[K, V]
	next  *Entry[K, V]
}

func (me *Entry[K, V]) append(entry *Entry[K, V]) {
	if me.next == nil {
		entry.prev = me
		me.next = entry
		return
	}

	me.next.append(entry)
}

func (me *Entry[K, V]) iter(fn func(me *Entry[K, V]) (stop bool)) {
	if stop := fn(me); stop {
		return
	}

	if me.next == nil {
		return
	}

	me.next.iter(fn)
}

func (me *Entry[K, V]) remove(key K) {
	if me.Key == key {
		me.prev = me.next
		return
	}

	me.next.remove(key)
}
