package matrix

import "gonum.org/v1/gonum/mat"

func areUniformedVecs(vecs []Vector) (int, bool) {
	if len(vecs) == 0 {
		return 0, true
	} else if len(vecs) == 1 {
		return len(vecs[0]), true
	}

	for ind := 1; ind < len(vecs); ind++ {
		if len(vecs[ind]) != len(vecs[ind -1 ]){
			return 0, false
		}
	}

	return len(vecs[0]), true
}

func DenseToMatrix(dense *mat.Dense) (Matrix, error) {
	width, length := dense.Dims()
	mat, err := NewEmptyMatrix(width)
	if err != nil {
		return Matrix{}, err
	}

	for c := 0; c < length; c++ {
		vec, err := NewVector(width)
		if err != nil {
			return Matrix{}, err
		}

		for r := 0; r < width; r++ {
			vec[r] = dense.At(r, c)
		}

		err = mat.AppendNewVec(vec)
		if err != nil {
			return Matrix{}, err
		}
	}
	return mat, nil
}