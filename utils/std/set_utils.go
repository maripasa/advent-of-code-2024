package std 

// ===== Set Utilities ===== setutils =====

// A set is a collection of unique elements.
type Set[T comparable] map[T]struct{}

// Adds an element to the set.
func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

// Removes an element from the set.
func (s Set[T]) Remove(element T) {
	delete(s, element)
}

// Checks if an element is in the set.
func (s Set[T]) Contains(element T) bool {
	_, exists := s[element]
	return exists
}

// Returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s)
}
