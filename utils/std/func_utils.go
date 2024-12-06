package std

// ===== Functional Utilities ===== funcutils =====

// Returns a new slice containing only the elements from the first slice for which the given function does not return an error.
func FilterMap[T any](slice []T, predicate func(T) error) []T {
	var result []T
	for _, value := range slice {
		if predicate(value) == nil {
			result = append(result, value)
		}
	}
	return result
}

// Returns a new slice containing only the elements that return true when passed through a predicate.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// Returns a value after collapsing the slice with an accumulator function.
func Reduce[T any, U any](slice []T, initial U, accumulator func(U, T) U) U {
	result := initial
	for _, value := range slice {
		result = accumulator(result, value)
	}
	return result
}

// Joins a slice of slices into a single slice.
func Flatten[T any](slice [][]T) []T {
	var result []T
	for _, outer := range slice {
		for _, inner := range outer {
			result = append(result, inner)
		}
	}
	return result
}

// Applies a transform to every value in a slice, returning a new slice.
func Map[T any, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, value := range slice {
		result[i] = transform(value)
	}
	return result
}
