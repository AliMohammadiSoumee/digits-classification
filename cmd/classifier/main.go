package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/alidadar7676/digits-classification/classifier"
	"github.com/alidadar7676/digits-classification/digit"
	"github.com/sirupsen/logrus"
)

func main() {
	c, err := classifier.NewClassifier("/home/ali/Developer/Go/src/github.com/alidadar7676/digits-classification/USPSdata/Numerals")
	if err != nil {
		logrus.Error(err)
		return
	}

	dir := "/home/ali/Developer/Go/src/github.com/alidadar7676/digits-classification/USPSdata/Test"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		logrus.Error(err)
		return
	}

//	temp := make(map[string]map[string]int)
//	temp["num0"] = make(map[string]int)
//	temp["num1"] = make(map[string]int)
//	temp["num2"] = make(map[string]int)
//	temp["num3"] = make(map[string]int)
//	temp["num4"] = make(map[string]int)
//	temp["num5"] = make(map[string]int)
//	temp["num6"] = make(map[string]int)
//	temp["num7"] = make(map[string]int)
//	temp["num8"] = make(map[string]int)
//	temp["num9"] = make(map[string]int)
//
	cnt := 0
	for _, file := range files {
		name := file.Name()[7:11]
		filePath := filepath.Join(dir, file.Name())
		digit, err := digit.NewDigit(filePath)
		if err != nil {
			logrus.Error(err)
			return
		}
		vec, err := digit.Vector()
		if err != nil {
			logrus.Error(err)
			return
		}

		ans := c.Classify((vec))
		if ans == name {
			cnt++
		}
		fmt.Println(name, ans)
		fmt.Println("count =", cnt)
	}

	//for key, val := range temp {
	//	fmt.Println(key, "------------------------------")
	//	for key1, val1 := range val {
	//		fmt.Println(key1, "--->> ", val1)
	//	}
	//}
	//vec1 := []float64{1, 2, 1}
	//vec2 := []float64{3, 5, 2}

	//m, err := matrix.NewMatrixFromVectors([]matrix.Vector{vec1, vec2})
	//if err != nil {
	//	println(err)
	//}
	//class.space = m
	//return class, nil

}
