package digit

import (
	"image"
	"math"
	"fmt"
	"image/color"
	"github.com/alidadar7676/digits-classification/matrix"
)

func forEachPixel(size image.Point, f func(x int, y int)) {
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			f(i, j)
		}
	}
}

func grayScale(img image.Image) *image.Gray {
	gray := image.NewGray(img.Bounds())
	size := img.Bounds().Size()
	forEachPixel(size, func(x, y int) {
		gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
	})
	return gray
}

func VectorToImage(vec matrix.Vector) {
	for i := 0; i < ImageHeigth; i++ {
		for j := 0; j < ImageWidth; j++ {
			num := 0
			if math.Abs(vec.At(i*ImageWidth + j)) >= 10000 {
				num = 1
			} else {
				num = 0
			}
			fmt.Print(fmt.Sprintf("%d ", num))
		}
		fmt.Println()
	}
}