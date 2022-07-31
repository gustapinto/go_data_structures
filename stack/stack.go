package stack

type StackNode struct {
	Value any
	Next  *StackNode
}

// Size retorna o tamanho atual da pilha
func Size(stack *StackNode) (count uint) {
	if stack == nil {
		return 0
	}

	temp := stack

	for temp != nil {
		temp = temp.Next
		count++
	}

	return count
}

// Push adiciona um novo elemento no topo da pilha
func Push(stack *StackNode, item any) *StackNode {
	if stack == nil {
		return &StackNode{Value: item}
	}

	temp := *stack

	*stack = StackNode{
		Value: item,
		Next:  &temp,
	}

	return stack
}

// Pop remove o último elemento da pilha
func Pop(stack *StackNode) any {
	if Size(stack) == 0 {
		return nil
	}

	if Size(stack) == 1 {
		value := stack.Value
		*stack = StackNode{}

		return value
	}

	value := stack.Value
	*stack = *stack.Next

	return value
}

// Traverse retorna uma slice com todos os valores na pilha
func Traverse(stack *StackNode) (values []any) {
	if stack == nil {
		return []any{}
	}

	temp := stack

	for temp != nil {
		values = append(values, temp.Value)
		temp = temp.Next
	}

	return values
}
