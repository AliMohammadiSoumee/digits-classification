package matrix

import "fmt"

// Vector is a slice of float64 values
type Vector []float64

// At returns the ind-th index of vector, starts from 0
func (v Vector) At(ind int) (float64, error) {
	if ind >= len(v) {
		return 0, fmt.Errorf("vector: Index out of range")
	}
	return v[ind], nil
}

// Len returns length of the vector
func (v Vector) Len() int {
	return len(v)
}

// NewVector creates a new vector with a specific length
func NewVector(l int) (Vector, error) {
	if l < 0 {
		return nil, fmt.Errorf("Vector: Length of vector must be equal or greater than zero")
	}
	return make(Vector, l), nil
}
