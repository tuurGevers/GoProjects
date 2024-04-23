package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// image with width and height
type Image struct {
	width, height int
}

// implement interface methods if image for struct Image
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.height, img.height)
}

func (img *Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func mainImg() {
	m := &Image{256, 256}
	pic.ShowImage(m)
}
