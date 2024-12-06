package std

// ===== Slice Utilities ===== sliceutils =====

// Counts all occurrences of a specific value in a slice.
func Count[T comparable](slice []T, value T) int {
	count := 0
	for _, item := range slice {
		if item == value {
			count++
		}
	}
	return count
}

// Checks if a value is present in a slice.
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Inserts an element at a specific index in a slice.
func InsertAtIndex[T any](slice []T, index int, value T) []T {
	if index < 0 || index > len(slice) {
		return slice
	}
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}

// Removes an element at a specific index in a slice.
func RemoveAtIndex[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func Split[T any](slice []T) ([]T, []T) {
	left, right := []T{}, []T{}
	for i, item := range slice {
		if i%2 == 0 {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}
	return left, right
}

func Zip[T any](left []T, right []T) [][2]T {
  result := make([][2]T, len(left))
  for i, item := range left {
    result[i] = [2]T{item, right[i]}
  }
  return result
}

func Unzip[T any] (slice [][2]T) ([]T, []T){
  left, right := make([]T, len(slice)), make([]T, len(slice))
  for i, item := range slice {
    left[i], right[i] = item[0], item[1]
  }
  return left, right
}
