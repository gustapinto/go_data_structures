package set

// Set (conjunto) é uma coleção que não permite itens duplicados, portanto a
// implementação mais prática é utilizando um map como container
type Set[T comparable] struct {
	Values map[T]bool
}

// Create atua como um alias para Push(nil, value)
func Create[T comparable](value T) *Set[T] {
	return Push(nil, value)
}

// Contains indica se um elemento já existe no conjunto
func Contains[T comparable](set *Set[T], value T) bool {
	if set == nil {
		return false
	}

	_, exists := set.Values[value]

	return exists
}

// Push adiciona um novo elemento no conjunto, apenas se ele não existir ainda
func Push[T comparable](set *Set[T], value T) *Set[T] {
	if set == nil {
		return &Set[T]{
			Values: map[T]bool{value: true},
		}
	}

	if set.Values == nil {
		set.Values = map[T]bool{}
	}

	if Contains(set, value) {
		return set
	}

	set.Values[value] = true

	return set
}

// Traverse retorna todos os valores contidos no set
func Traverse[T comparable](set *Set[T]) (values []T) {
	if set == nil {
		return []T{}
	}

	for value, _ := range set.Values {
		values = append(values, value)
	}

	return values
}

// Union retorna uma novo conjunto com todos os valores de dois conjuntos
func Union[T comparable](set1, set2 *Set[T]) (union *Set[T]) {
	if set1 == nil && set2 == nil {
		return nil
	} else if set1 == nil && set2 != nil {
		return set2
	} else if set1 != nil && set2 == nil {
		return set1
	}

	for _, value := range Traverse(set1) {
		union = Push(union, value)
	}
	for _, value := range Traverse(set2) {
		union = Push(union, value)
	}

	return union
}

// Intersection retorna um novo conjunto com todo os valores que formam a
// intersecção entre dois conjuntos
func Intersection[T comparable](set1, set2 *Set[T]) (intersection *Set[T]) {
	if set1 == nil || set2 == nil {
		return nil
	}

	for _, value := range Traverse(set1) {
		if Contains(set2, value) {
			intersection = Push(intersection, value)
		}
	}

	return intersection
}

// Difference retorna um novo conjunto com todos os valores do primeiro conjunto
// que não estão no segundo
func Difference[T comparable](set1, set2 *Set[T]) (difference *Set[T]) {
	if set1 == nil {
		return nil
	} else if set1 != nil && set2 == nil {
		return set1
	}

	for _, value := range Traverse(set1) {
		if !Contains(set2, value) {
			difference = Push(difference, value)
		}
	}

	return difference
}
