package matrix

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Matrix is 2D slice of float64 values
type Matrix [][]float64

// Dims returns dimension(size) of the matrix
func (m Matrix) Dims() (r, c int) {
	if len(m) == 0 {
		return 0, 0
	}
	return len(m), len(m[0])
}

// At return stored value in position (row, col) of the matrix
func (m Matrix) At(row, col int) float64 {
	return m[row][col]
}

// T returns the transpose of the matrix
func (m Matrix) T() mat.Matrix {
	row, col := m.Dims()

	mt, _ := NewMatrix(col, row)
	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			mt[c][r] = m[r][c]
		}
	}
	return mt
}

// Factorize returns the singular value decomposition of the matrix
func (m Matrix) Factorize() *mat.SVD {
	svd := mat.SVD{}
	svd.Factorize(m, mat.SVDFull)

	return &svd
}

// NewMatrixFromVectors Creates a new matrix which each vector is stored vertically in it
func NewMatrixFromVectors(vecs []Vector) (Matrix, error) {
	col := len(vecs)
	if col <= 0 {
		return Matrix{}, fmt.Errorf("Matrix: Empty set of vectors")
	}
	row := vecs[0].Len()

	mat, err := NewMatrix(row, col)
	if err != nil {
		return Matrix{}, err
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			val, err := vecs[c].At(r)
			if err != nil {
				return Matrix{}, err
			}
			mat[r][c] = val
		}
	}

	return mat, nil
}

// NewMatrix creates a new matrix with a specific dimension
func NewMatrix(row, col int) (Matrix, error) {
	if row < 0 {
		return nil, fmt.Errorf("Matrix: Length of rows must be greater than zero")
	}

	if col < 0 {
		return nil, fmt.Errorf("Matrix: Length of columns must be greater than zero")
	}

	mat := make([][]float64, row)
	for r := 0; r < row; r++ {
		mat[r] = make([]float64, col)
	}

	return mat, nil
}
