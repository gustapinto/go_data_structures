package doublylinkedlist

type ListNode[T comparable] struct {
	Previous *ListNode[T]
	Value    T
	Next     *ListNode[T]
}

func AddNode[T comparable](node *ListNode[T], value T) *ListNode[T] {
	if node == nil {
		return &ListNode[T]{Value: value}
	}

	if value == node.Value {
		return node
	}

	if node.Next == nil {
		temp := node

		node.Next = &ListNode[T]{
			Value:    value,
			Previous: temp,
		}

		return node.Next
	}

	return AddNode(node.Next, value)
}

// Traverse retorna uma slice com todos os elementos da lista duplamente ligada, em ordem
func Traverse[T comparable](node *ListNode[T]) (values []T) {
	if node == nil {
		return []T{}
	}

	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}

	return values
}

// Reverse retorna uma slice com todos os elementos da lista em ordem reversa
func Reverse[T comparable](node *ListNode[T]) (values []T) {
	if node == nil {
		return []T{}
	}

	// Cria uma variável temporária
	temp := node

	for node != nil {
		temp = node
		node = node.Next
	}

	for temp.Previous != nil {
		values = append(values, temp.Value)
		temp = temp.Previous
	}

	values = append(values, temp.Value)

	return values
}

func Size[T comparable](node *ListNode[T]) int {
	if node == nil {
		return 0
	}

	counter := 0
	for node != nil {
		counter++

		node = node.Next
	}

	return counter
}

func Lookup[T comparable](node *ListNode[T], value T) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if node.Next == nil {
		return false
	}

	return Lookup(node, value)
}
