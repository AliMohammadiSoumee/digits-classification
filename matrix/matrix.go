package matrix

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"github.com/sirupsen/logrus"
)

// Matrix is 2D slice of float64 values
type Matrix struct {
	mat []Vector
	length, width int
}

// Dims returns dimension(size) of the matrix
func (m Matrix) Dims() (r, c int) {
	return m.length, m.width
}

// At return stored value in position (row, col) of the matrix
func (m Matrix) At(row, col int) float64 {
	return m.mat[col].At(row)
}

// T returns the transpose of the matrix
func (m Matrix) T() mat.Matrix {
	mtrx, err := NewEmptyMatrix(m.length)
	if err != nil {
		logrus.Error("matrix: Cannot create matrix transpose")
		return Matrix{}
	}

	for r := 0; r < m.width; r++ {
		vec, err := NewVector(m.length)
		if err != nil {
			return Matrix{}
		}
		for c := 0; c < m.length; c++ {
			vec[c] = m.At(r, c)
		}

		mtrx.AddVectorOnCol(vec, r)
	}

	return mtrx
}

// Factorize returns the singular value decomposition of the matrix
func (m Matrix) Factorize() *mat.SVD {
	svd := mat.SVD{}
	svd.Factorize(m, mat.SVDFull)

	return &svd
}

// AddVectorOnCol puts a vector on the column col of the matrix
func (m *Matrix) AddVectorOnCol(vec Vector, col int) error {
	if vec.Len() != m.width {
		return fmt.Errorf("Matrix: Vector's width is out of range")
	}

	m.mat[col] = vec
	return nil
}

func (m *Matrix) AppendNewVec(vec Vector) error {
	if vec.Len() != m.width {
		return fmt.Errorf("Matrix: Vector's width is out of range")
	}

	m.mat = append(m.mat, vec)
	m.length++
	return nil
}

// NewEmptyMatrix returns EmptyMatrix
func NewEmptyMatrix(width int) (Matrix, error) {
	return Matrix{
		mat: make([]Vector, 0),
		length: 0,
		width: width,
	}, nil
}

// NewMatrixFromVectors Creates a new matrix which each vector is stored vertically in it
func NewMatrixFromVectors(vecs []Vector) (Matrix, error) {
	width, ok := areUniformedVecs(vecs)
	if !ok {
		return Matrix{}, fmt.Errorf("matrix: Vectors don't have a same length")
	}

	mtrx, err := NewEmptyMatrix(width)
	if err != nil {
		return Matrix{}, err
	}

	for _, vec := range vecs {
		mtrx.AddVectorOnCol(vec, mtrx.length)
	}

	return mtrx, nil
}

// NewMatrix creates a new matrix with a specific dimension
func NewMatrix(length, width int) (Matrix, error) {
	if length < 0 {
		return Matrix{}, fmt.Errorf("Matrix: Length of matrix must be greater than zero")
	}

	if width < 0 {
		return Matrix{}, fmt.Errorf("Matrix: Length of matrix must be greater than zero")
	}

	mtrx := Matrix{
		mat: make([]Vector, 0),
		length: 0,
		width: width,
	}

	for i := 0; i < length; i++ {
		vec, err := NewVector(width)
		if err != nil {
			return Matrix{},err
		}
		err = mtrx.AddVectorOnCol(vec, mtrx.length)
		if err != nil {
			return Matrix{}, err
		}
		mtrx.length++
	}

	return mtrx , nil
}