package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

/**
Write another one, but this time it will return an implementation of image.Image instead of a slice of data.
Define your own Image type, implement the necessary methods, and call pic.ShowImage.
Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
ColorModel should return color.RGBAModel. At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
**/

type Image struct {
	x int
	y int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{256, 65}
	pic.ShowImage(m)
}
