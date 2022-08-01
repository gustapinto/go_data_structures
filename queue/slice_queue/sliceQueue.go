package slicequeue

// Implementa uma fila usando uma slice como container, ao invés de uma
// lista ligada
type SliceQueue struct {
	Values []any
}

// Size retorna o tamanho atual da fila
func Size(queue *SliceQueue) uint {
	if queue == nil {
		return 0
	}

	return uint(len(queue.Values))
}

// Push adiciona um novo elemento no final da fila
func Push(queue *SliceQueue, value any) *SliceQueue {
	if queue == nil {
		return &SliceQueue{
			Values: []any{value},
		}
	}

	queue.Values = append(queue.Values, value)

	return queue
}

// Pop remove o primeiro elemento da fila
func Pop(queue *SliceQueue) any {
	if Size(queue) == 0 {
		return []any{}
	}

	value := queue.Values[0]
	// Remove o primeiro elemento da fila redefinando a slice de valores
	// para receber todos os seus elementos menos o primeiro
	queue.Values = queue.Values[1:]

	return value
}

// Traverse retorna todos os elementos da fila
func Traverse(queue *SliceQueue) []any {
	if Size(queue) == 0 {
		return []any{}
	}

	return queue.Values
}
