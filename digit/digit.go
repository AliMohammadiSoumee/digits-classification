package digit

import (
	"image"
	"image/png"
	"os"

	"github.com/alidadar7676/digits-classification/matrix"
	"github.com/nfnt/resize"
)

type Digit struct {
	img *image.Gray
	vec matrix.Vector
}

func (d *Digit) Vector() (matrix.Vector, error) {
	if d.vec != nil {
		return d.vec, nil
	}

	dims := d.img.Bounds().Size()
	vec, err := matrix.NewVector(dims.X * dims.Y)
	if err != nil {
		return nil, err
	}

	for col := 0; col < dims.X; col++ {
		for row := 0; row < dims.Y; row++ {
			vec = append(vec, float64(d.img.GrayAt(row, col).Y))
		}
	}
	d.vec = vec

	return d.vec, nil
}

func NewDigit(path string) (Digit, error) {
	file, err := os.Open(path)
	if err != nil {
		return Digit{}, err
	}
	defer file.Close()

	file.Seek(0, 0)

	img, err := png.Decode(file)
	if err != nil {
		return Digit{}, err
	}
	img16x16 := resize.Resize(16, 16, img, resize.NearestNeighbor)

	grayImg := grayScale(img16x16)
	

	return Digit{
		img: grayImg,
	}, nil
}
