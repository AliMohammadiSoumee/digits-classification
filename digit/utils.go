package digit

import (
	"image"
	"image/color"
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
