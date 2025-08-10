package orderedmap

type mapEntry[K comparable, V any] struct {
	Key   K
	Value V
	Next  *mapEntry[K, V]
}

func (me *mapEntry[K, V]) append(entry *mapEntry[K, V]) {
	if me.Next == nil {
		me.Next = entry
		return
	}

	me.Next.append(entry)
}

func (me *mapEntry[K, V]) traverse(fn func(me *mapEntry[K, V]) (stop bool)) {
	if me.Next == nil {
		return
	}

	if stop := fn(me); stop {
		return
	}

	me.Next.traverse(fn)
}

func (me *mapEntry[K, V]) search(key K) (value V, exists bool) {
	if me.Key == key {
		return me.Value, true
	}

	if me.Next == nil {
		return value, false
	}

	return me.Next.search(key)
}
