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

func (m Matrix) VectorAtIndex(ind int) Vector {
	return m.mat[ind]
}

func (m Matrix) String() string {
	str := ""
	str += fmt.Sprintf("%d ---- %d\n", m.width, m.length)
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.length; j++ {
			str += fmt.Sprintf("%f ", m.At(i, j))
		}
		str += fmt.Sprintf("\n")
	}
	return str
}

// Dims returns dimension(size) of the matrix
func (m Matrix) Dims() (r, c int) {
	return m.width, m.length
}

// At return stored value in position (row, col) of the matrix
func (m Matrix) At(row, col int) float64 {
	return m.mat[col].At(row)
}

// T returns the transpose of the matrix
func (m Matrix) T() mat.Matrix {
	mtrx, err := NewEmptyMatrix(m.length)
	if err != nil {
		logrus.Error(err)
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

		err = mtrx.AppendNewVec(vec)
		if err != nil {
			logrus.Error(err)
			return Matrix{}
		}
	}

	return mtrx
}

// Factorize returns the singular value decomposition of the matrix
func (m Matrix) Factorize() *mat.SVD {
	svd := mat.SVD{}
	svd.Factorize(m, mat.SVDThin)

	return &svd
}

func (m Matrix) MultiplyVector(vec Vector) (Vector, error) {
	if m.length != vec.Len() {
		return Vector{}, fmt.Errorf("matrix: matrix and vector don't have a same dimensions")
	}

	mulVec, err := NewVector(vec.Len())
	if err != nil {
		return Vector{}, err
	}

	for r := 0; r < m.width; r++ {
		val := 0.0
		for c := 0; c < m.length; c++ {
			val += m.At(r, c) * vec.At(r)
		}
		mulVec[r] = val
	}

	return mulVec, nil
}

func (m Matrix) MultiplyMatrix(n Matrix) (Matrix, error) {
	c := &mat.Dense{}
	c.Mul(m, n)
	res, err := DenseToMatrix(c)
	if err != nil {
		return Matrix{}, err
	}
	return res, nil
}

// AddVectorOnCol puts a vector on the column col of the matrix
func (m *Matrix) AddVectorOnCol(vec Vector, col int) error {
	if vec.Len() != m.width || col >= m.length {
		return fmt.Errorf("Matrix: Index out of range")
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

func (m *Matrix) SplitCol(num int) (Matrix, error) {
	if m.length <= num {
		return Matrix{}, fmt.Errorf("Matrix: num for split is greater than the length of matrix")
	}

	mat, err := NewEmptyMatrix(m.width)
	if err != nil {
		return Matrix{}, err
	}

	for j := 0; j < num; j++ {
		vec, err := NewVector(m.width)
		if err != nil {
			return Matrix{}, err
		}

		for i := 0; i < m.width; i++ {
			vec[i] = m.At(i, j)
		}
		err = mat.AppendNewVec(vec)
		if err != nil {
			return Matrix{}, err
		}
	}

	return mat, nil
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
		err := mtrx.AppendNewVec(vec)
		if err != nil {
			return Matrix{}, err
		}
	}

	return mtrx, nil
}

// NewMatrix creates a new matrix with a specific dimension
func NewMatrix(width, length int) (Matrix, error) {
	if length < 0 {
		return Matrix{}, fmt.Errorf("Matrix: Length of matrix must be greater than zero")
	}

	mtrx, err := NewEmptyMatrix(width)
	if err != nil {
		return Matrix{}, nil
	}

	for i := 0; i < length; i++ {
		vec, err := NewVector(width)
		if err != nil {
			return Matrix{},err
		}
		err = mtrx.AppendNewVec(vec)
		if err != nil {
			return Matrix{}, err
		}
	}

	return mtrx , nil
}