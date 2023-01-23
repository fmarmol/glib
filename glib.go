package glib

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type Image struct {
	indexRef int // for subimage
	stride   int // for subimage
	width    int
	height   int
	pixels   []Pixel
}

func NewImageFromImage(img image.Image) *Image {
	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	ret := NewImage(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			ret.Set(x, y, img.At(x, y))
		}
	}
	return ret
}

func NewImage(w, h int) *Image {
	return &Image{
		pixels: make([]Pixel, w*h, w*h),
		width:  w,
		height: h,
		stride: w,
	}
}

func (i *Image) Set(x, y int, c color.Color) {
	index := i.indexRef + y*i.stride + x
	r, g, b, a := c.RGBA()
	i.pixels[index].SetR(uint8(r))
	i.pixels[index].SetG(uint8(g))
	i.pixels[index].SetB(uint8(b))
	i.pixels[index].SetA(uint8(a))
}

func (i *Image) Fill(c color.Color) {
	for x := 0; x < i.width; x++ {
		for y := 0; y < i.height; y++ {
			i.Set(x, y, c)
		}
	}
}

func (i *Image) SubImage(x, y, w, h int) *Image {
	r := &Image{
		pixels:   i.pixels,
		width:    w,
		height:   h,
		stride:   i.stride,
		indexRef: i.indexRef + y*i.stride + x,
	}
	return r
}

// implement image.Image

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i *Image) ColorModel() color.Model {
	return color.NRGBAModel
}

func (i *Image) At(x, y int) color.Color {
	index := i.indexRef + y*i.stride + x
	return i.pixels[index]
}

func (i *Image) ToPng(w io.Writer) error {
	return png.Encode(w, i)
}
