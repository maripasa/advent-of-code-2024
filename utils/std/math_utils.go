package std 

// ===== Math Utilities ===== mathutils =====

// Calculates the factorial of a number.
func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Returns the absolute value of an integer.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Represents a Number.
type Number interface {
  ~int | ~float64 | ~float32
}

// Represents a Vector in 2d space.
type Vector2[T Number] struct {
  X, Y T 
}

// Adds a Vector2[T] to a Vector2[T]
func (v Vector2[T]) Add(v2 Vector2[T]) (Vector2[T]) {
  return Vector2[T]{ X: v.X + v2.X, Y: v.Y + v2.Y }
}

// Subtracts a Vector2[T] from a Vector2[T]
func (v Vector2[T]) Subtract(v2 Vector2[T]) (Vector2[T]) {
  return Vector2[T]{ X: v.X - v2.X, Y: v.Y - v2.Y }

}// Scale a vector by a scalar
func (v Vector2[T]) Scale(scalar T) Vector2[T] {
	return Vector2[T]{ X: v.X * scalar, Y: v.Y * scalar }
}

// Dot product of two vectors
func (v Vector2[T]) Dot(v2 Vector2[T]) T {
	return v.X*v2.X + v.Y*v2.Y
}

