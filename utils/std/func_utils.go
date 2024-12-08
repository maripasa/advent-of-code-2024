package std

// ===== Functional Utilities ===== funcutils =====

// Returns a new slice containing only the elements from the first slice for which the given function does not return an error.
func FilterMap[T any, U any](slice []T, predicate func(T) (U, error)) []U {
	var result []U
	for _, value := range slice {
    if answer, err := predicate(value); err == nil {
			result = append(result, answer)
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

// Counts the number of elements in a given list satisfying a given preicate
func Count[T any](slice []T, predicate func(T) bool) int {
  if len(slice) == 0 {
    return 0
  }

  if predicate(slice[0]) {
    return 1 + Count(slice[1:], predicate)
  }
  return Count(slice[1:], predicate)
}

// Returns true if the given function returns true for all the elements in the given list. If the function returns false for any of the elements it immediately returns false without checking the rest of the list.
func All[T any](list []T, predicate func(T) bool) bool {
	for _, elem := range list {
		if !predicate(elem) {
			return false
		}
	}
	return true
}
