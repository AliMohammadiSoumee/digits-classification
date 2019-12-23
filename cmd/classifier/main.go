package main

import (
	"fmt"
	"github.com/alidadar7676/digits-classification/matrix"
)

func main() {
	mat, err := matrix.NewMatrix(2, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	mat[0][0] = 1
	mat[0][1] = 2
	mat[1][0] = 3
	mat[1][1] = 5
	fmt.Println(mat.Factorize())

	vec1 := matrix.Vector{1, 3}
	vec2 := matrix.Vector{2, 5}
	matt, err := matrix.NewMatrixFromVectors([]matrix.Vector{vec1, vec2})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(matt.Factorize())
}