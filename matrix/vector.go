package matrix

import "fmt"

import "github.com/sirupsen/logrus"

// Vector is a slice of float64 values
type Vector []float64

// At returns the ind-th index of vector, starts from 0
func (v Vector) At(ind int) float64 {
	if ind >= len(v) || ind < 0 {
		logrus.Error(fmt.Printf("vector: Index out of range. len = %i, ind = %i", v.Len(), ind))
		return 0
	}
	return v[ind]
}

// Len returns length of the vector
func (v Vector) Len() int {
	return len(v)
}

// NewVector creates a new vector with a specific length
func NewVector(l int) (Vector, error) {
	if l < 0 {
		return nil, fmt.Errorf("vector: Length of vector must be equal or greater than zero. len = %i", l)
	}
	return make(Vector, l), nil
}
