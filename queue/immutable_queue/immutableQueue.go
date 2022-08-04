package immutablequeue

import "errors"

var ErrCannotPopOnEmptyQueue = errors.New("cannot use pop on a empty queue")

// Implementa uma fila imutável, na qual as operações não modificm seus valores,
// mas sim retornam novas filas com os novos valores
type ImmutableQueue[T any] struct {
	Values []T
}

// NewImmutableQueue cria uma nova fila com os valores passados, em ordem
func NewImmutableQueue[T any](values ...T) ImmutableQueue[T] {
	return ImmutableQueue[T]{
		Values: values,
	}
}

// Size retorna o tamanho da fila
func Size[T any](queue ImmutableQueue[T]) uint {
	return uint(len(queue.Values))
}

// Push adiciona um novo elemento no final da fila
func Push[T any](queue ImmutableQueue[T], values ...T) ImmutableQueue[T] {
	return NewImmutableQueue(append(queue.Values, values...)...)
}

// Pop remove um elemento do começo da fila, retornando o elemento e uma
// fila sem esse elemento
func Pop[T any](queue ImmutableQueue[T]) (T, ImmutableQueue[T], error) {
	if Size(queue) == 0 {
		var zeroValue T
		var zeroQueue ImmutableQueue[T]

		return zeroValue, zeroQueue, ErrCannotPopOnEmptyQueue
	}

	if Size(queue) == 1 {
		return queue.Values[0], ImmutableQueue[T]{Values: []T{}}, nil
	}

	return queue.Values[0], NewImmutableQueue(queue.Values[1:]...), nil
}

// Traverse retorna todos os elementos da fila
func Traverse[T any](queue ImmutableQueue[T]) []T {
	if Size(queue) == 0 {
		return []T{}
	}

	return queue.Values
}
