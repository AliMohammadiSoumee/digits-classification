package matrix

import "fmt"

import "github.com/sirupsen/logrus"

import "math"

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

func (v Vector) Minus(vec Vector) (Vector, error) {
	if v.Len() != vec.Len() {
		return Vector{}, fmt.Errorf("vector: Can't Minus two vector with different sizes")
	}

	minVec, err := NewVector(v.Len())
	if err != nil {
		return Vector{}, err
	}

	for i := range v {
		minVec[i] = v[i] - vec[i]
	}

	return minVec, nil
}

func (v Vector) Norm2() float64 {
	var val float64 = 0.0
	for _, i := range v {
		val += (i * i)
		//fmt.Println(i, i * i, "----->>>", val)
	}
	//fmt.Println(val, math.Sqrt(val))
	return math.Sqrt(val)
}

// NewVector creates a new vector with a specific length
func NewVector(l int) (Vector, error) {
	if l < 0 {
		return nil, fmt.Errorf("vector: Length of vector must be equal or greater than zero. len = %i", l)
	}
	return make(Vector, l), nil
}
