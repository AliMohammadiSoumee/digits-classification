package class

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/alidadar7676/digits-classification/digit"
	"github.com/alidadar7676/digits-classification/matrix"
	"github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/mat"
)

type Class struct {
	baseDir    string
	space      matrix.Matrix
	identifier matrix.Matrix
}

func (c *Class) addDigitToSpace(path string) error {
	if !isPng(path) {
		return fmt.Errorf("class: File %s is not png", path)
	}

	digit, err := digit.NewDigit(path)
	if err != nil {
		return fmt.Errorf("class: Error when creating new digit with path = %s: %s", path, err)
	}

	vec, err := digit.Vector()
	if err != nil {
		return err
	}
	if err := c.space.AppendNewVec(vec); err != nil {
		return err
	}

	return nil
}

func (c *Class) Distance(vec matrix.Vector) float64 {
	var infinity float64 = 100000000000

	projVec, err := c.identifier.MultiplyVector(vec)
	if err != nil {
		logrus.Error(err)
		return infinity
	}

	disVec, err := vec.Minus(projVec)
	if err != nil {
		logrus.Error(err)
		return infinity
	}

	norm := disVec.Norm2()
	return norm
}

func (c *Class) makeIdentifier() error {
	//vec1 := []float64{1, 2 , 1, 4}
	//vec2 := []float64{3, 5 , 2, 3}
	//vec3 := []float64{7, 4, 2, 9}

	//m, err := matrix.NewMatrixFromVectors([]matrix.Vector{vec1, vec2, vec3})
	//if err != nil {
	//	logrus.Error(err)
	//}

	//svd := m.Factorize()

	//logrus.Info("class: factorize space")
	//u := &mat.Dense{}
	//svd.UTo(u)

	//uMat, err := matrix.DenseToMatrix(u)
	//if err != nil {
	//	return err
	//}
	//uMatT := uMat.T().(matrix.Matrix)

	//fmt.Println("=======================================")
	//fmt.Println("===== Matrix space")
	//fmt.Println(m)
	//fmt.Println("===== Matrix U")
	//fmt.Println(uMat)
	//fmt.Println("===== Matrix U.T()")
	//fmt.Println(uMat.T())
	//fmt.Println("===== Matrix U * U.T()")
	//fmt.Println(uMat.MultiplyMatrix(uMatT))
	//fmt.Println("===== Matrix U.T() * U")
	//fmt.Println(uMatT.MultiplyMatrix(uMat))
	//fmt.Println("---------------------------------------")

	logrus.Info("Class: Factorize space")
	svd := c.space.Factorize()

	u := &mat.Dense{}
	svd.UTo(u)
	uMat, err := matrix.DenseToMatrix(u)
	if err != nil {
		return err
	}

	uMat, err = uMat.SplitCol(2)
	if err != nil {
		return err
	}

	//	fmt.Println("===== Matrix U")
	//	fmt.Println(uMat)
	//	uMatT := uMat.T().(matrix.Matrix)
	//	fmt.Println("===== Matrix U * U.T()")
	//	fmt.Println(uMat.MultiplyMatrix(uMatT))
	//	fmt.Println("===== Matrix U.T() * U")
	//	fmt.Println(uMatT.MultiplyMatrix(uMat))
	//

	logrus.Info("Class: Create identifier")
	c.identifier, err = uMat.MultiplyMatrix(uMat.T().(matrix.Matrix))
	if err != nil {
		return err
	}

	//for i := 0; i < 512; i++ {
	//	fmt.Print(c.identifier.At(0, i), "  ")
	//}
	//fmt.Println()
	//for i := 0; i < 512; i++ {
	//	fmt.Print(c.identifier.At(1, i), "  ")
	//}
	//fmt.Println()
	//fmt.Println("$$$$$$$$$$$$$$$")
	return nil
}

func NewClass(dir, name string) (Class, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return Class{}, err
	}

	space, err := matrix.NewEmptyMatrix(digit.ImageHeigth * digit.ImageWidth)
	if err != nil {
		return Class{}, err
	}

	class := Class{
		space:   space,
		baseDir: dir,
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if !strings.Contains(filePath, name) {
			continue
		}
		if err := class.addDigitToSpace(filePath); err != nil {
			logrus.Info("Class: Cannot add digit to space: ", err)
		}
	}
	logrus.Info("class: End Adding Images")

	if err := class.makeIdentifier(); err != nil {
		return Class{}, err
	}
	return class, nil
}
