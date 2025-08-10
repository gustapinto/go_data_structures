package orderedmap

import "sync"

type Map[K comparable, V any] struct {
	size    uint
	entries *mapEntry[K, V] // TODO: Trocar para implementação orimizada para busca, como árvore binária
	mu      *sync.Mutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		size:    0,
		entries: nil,
		mu:      &sync.Mutex{},
	}
}

func (om *Map[K, V]) Len() uint {
	return om.size
}

func (om *Map[K, V]) Set(key K, value V) {
	om.mu.Lock()
	defer om.mu.Unlock()

	entry := &mapEntry[K, V]{
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
	om.mu.Lock()
	defer om.mu.Unlock()

	if om.entries == nil {
		return value, false
	}

	return om.entries.search(key)
}

func (om *Map[K, V]) Iter(fn func(key K, value V) (stop bool)) {
	om.mu.Lock()
	defer om.mu.Unlock()

	if om.entries == nil {
		return
	}

	om.entries.traverse(func(me *mapEntry[K, V]) bool {
		return fn(me.Key, me.Value)
	})
}
