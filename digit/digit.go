package digit

import (
	"image"
	"image/png"
	"os"
	"strconv"

	"github.com/alidadar7676/digits-classification/matrix"
	"github.com/sirupsen/logrus"
)

const ImageWidth = 28
const ImageHeigth = 28

var maxW = 0
var maxH = 0

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
			vec[dims.Y*col+row] = float64(d.img.GrayAt(row, col).Y)
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

	//resizedImg := resize.Resize(ImageWidth, ImageHeigth, img, resize.Bicubic)
	//encode(resizedImg)
	
	grayImg := grayScale(img)

	//	if img.Bounds().Size().X > maxW {
	//		maxW = img.Bounds().Size().X
	//	}
	//	if img.Bounds().Size().Y > maxH {
	//		maxH = img.Bounds().Size().Y
	//	}
	//	fmt.Println(maxW, maxH)

	return Digit{
		img: grayImg,
	}, nil
}

var n int

func encode(img image.Image) {
	n++
	outputFile, err := os.Create("/home/ali/Developer/Go/src/github.com/alidadar7676/digits-classification/USPSdata/Test1/" + strconv.Itoa(n))
	if err != nil {
		logrus.Error(err)
		// Handle error
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, img)

	// Don't forget to close files
	outputFile.Close()
}
