package linkedlist

type ListNode[T comparable] struct {
	Value T
	Next  *ListNode[T]
}

// AddNode adiciona um novo nó na lista
func AddNode[T comparable](node *ListNode[T], value T) *ListNode[T] {
	if node == nil {
		return &ListNode[T]{Value: value}
	}

	if value == node.Value {
		return node
	}

	if node.Next == nil {
		node.Next = &ListNode[T]{Value: value}
	}

	return AddNode(node.Next, value)
}

// Traverse percorre todos os nós da lista, retornando uma slice com seus valores
func Traverse[T comparable](node *ListNode[T]) (values []T) {
	if node == nil {
		return []T{}
	}

	traverse := func(n *ListNode[T]) {
		for n != nil {
			values = append(values, n.Value)
			n = n.Next
		}
	}

	traverse(node)

	return values
}

// Lookup verifica se um elemento existe na lista
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

// Size retorna a contagem de todos os nós na lista
func Size[T comparable](node *ListNode[T]) int {
	if node == nil {
		return 0
	}

	count := 0
	for node != nil {
		count++

		node = node.Next
	}

	return count
}
