package class

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"gonum.org/v1/gonum/mat"

	"github.com/alidadar7676/digits-classification/digit"
	"github.com/sirupsen/logrus"
	"github.com/alidadar7676/digits-classification/matrix"
)

type Class struct {
	baseDir string
	space   matrix.Matrix
	svd mat.SVD
}

func (c *Class) createSpace() {
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

func NewClass(dir string) (Class, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return Class{}, err
	}

	space, err := matrix.NewEmptyMatrix(512)
	if err != nil {
		return Class{}, err
	}

	class := Class{
		space: space,
		baseDir: dir,
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if err := class.addDigitToSpace(filePath); err != nil {
			logrus.Info("Class: Cannot add digit to space: ", err)
		}
	}

	return class, nil
}
