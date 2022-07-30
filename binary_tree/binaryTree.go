package binary_tree

type Orderable interface {
	int | float64 | ~string
}

// BinaryTree implementa uma árvore binária em Go utilizando generics
type BinaryTree[T Orderable] struct {
	LeftNode  *BinaryTree[T] // Usa a própria definição da árvore para os galhos
	Value     T              // O valor do galho (Node)
	RightNode *BinaryTree[T]
}

// Insert adiciona um novo elemento na árvore, colocando-o na posição correta
func Insert[T Orderable](tree *BinaryTree[T], value T) *BinaryTree[T] {
	if tree == nil {
		return &BinaryTree[T]{
			Value: value,
		}
	}

	if value == tree.Value {
		return tree
	}

	if value < tree.Value {
		tree.LeftNode = Insert(tree.LeftNode, value)
	} else {
		tree.RightNode = Insert(tree.RightNode, value)
	}

	return tree
}

// Traverse opera percorrendo todos os galhos da árvore de forma recursiva, em order
// retornando uma slice de valores ordenados conforme a ordem em que foram atravessados
func Traverse[T Orderable](tree *BinaryTree[T]) (values []T) {
	// Usa uma função anônima interna para gerar uma slice sem precisar injetá-la como argumento
	var traverse func(t *BinaryTree[T])
	traverse = func(t *BinaryTree[T]) {
		if t == nil {
			return
		}

		traverse(t.LeftNode)
		values = append(values, t.Value)
		traverse(t.RightNode)
	}

	traverse(tree)

	return values
}

// TraversePostOrder percorre os galhos da árvore binária de forma recursiva e em pós-ordem
func TraversePostOrder[T Orderable](tree *BinaryTree[T]) (values []T) {
	var traverse func(t *BinaryTree[T])
	traverse = func(t *BinaryTree[T]) {
		if t == nil {
			return
		}

		traverse(t.LeftNode)
		traverse(t.RightNode)

		values = append(values, t.Value)
	}

	traverse(tree)

	return values
}

// TraversePostOrder percorre os galhos da árvore binária de forma recursiva e em pré-ordem
func TraversePreOrder[T Orderable](tree *BinaryTree[T]) (values []T) {
	var traverse func(t *BinaryTree[T])
	traverse = func(t *BinaryTree[T]) {
		if t == nil {
			return
		}

		values = append(values, t.Value)

		traverse(t.LeftNode)
		traverse(t.RightNode)
	}

	traverse(tree)

	return values
}
