package queue

type QueueNode struct {
	Value any
	Next  *QueueNode
}

// NewQueue atua como um alias para Push(nil, valor)
func NewQueue(value any) *QueueNode {
	return Push(nil, value)
}

// Size retorna o tamanho da fila
func Size(queue *QueueNode) uint {
	if queue == nil {
		return 0
	}

	temp := queue
	counter := uint(0)

	for temp != nil {
		temp = temp.Next
		counter++
	}

	return counter
}

// Push adiciona um novo item no final da fila
func Push(queue *QueueNode, item any) *QueueNode {
	if queue == nil {
		return &QueueNode{Value: item}
	}

	if queue.Next == nil {
		queue.Next = &QueueNode{Value: item}

		return queue.Next
	}

	return Push(queue.Next, item)
}

// Pop removoe um elemento do começo da fila
func Pop(queue *QueueNode) any {
	if Size(queue) == 0 {
		return nil
	}

	if Size(queue) == 1 {
		value := queue.Value
		// Quando a fila possui somente um item ao removê-lo a fila deve ser anulada, ou seja
		// a fila deve seu "zero value"
		*queue = QueueNode{}

		return value
	}

	value := queue.Value
	// Movimenta os elementos da fila quando o primeiro valor é removido
	*queue = *queue.Next

	return value
}

// Traverse retorna uma slice com todos os valores na fila
func Traverse(queue *QueueNode) (values []any) {
	if Size(queue) == 0 {
		return []any{}
	}

	temp := queue

	for temp != nil {
		values = append(values, temp.Value)
		temp = temp.Next
	}

	return values
}
