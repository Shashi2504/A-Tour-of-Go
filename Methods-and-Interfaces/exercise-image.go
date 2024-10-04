package main
import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i iamge) At(x, y int) color.Color{
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
