package glib

import (
	"image"
	"image/color"
)

type Pixel struct{ color.NRGBA }

func (p Pixel) RGBA() (r, g, b, a uint32) {
	return p.NRGBA.RGBA()
}

//go:inline
func (p *Pixel) R() uint8 {
	return p.NRGBA.R
}

//go:inline
func (p *Pixel) G() uint8 {
	return p.NRGBA.G
}

//go:inline
func (p *Pixel) B() uint8 {
	return p.NRGBA.B
}

//go:inline
func (p *Pixel) A() uint8 {
	return p.NRGBA.A
}

//go:inline
func (p *Pixel) SetR(v uint8) *Pixel {
	p.NRGBA.R = v
	return p
}

//go:inline
func (p *Pixel) SetG(v uint8) *Pixel {
	p.NRGBA.G = v
	return p
}

//go:inline
func (p *Pixel) SetB(v uint8) *Pixel {
	p.NRGBA.B = v
	return p
}

//go:inline
func (p *Pixel) SetA(v uint8) *Pixel {
	p.NRGBA.A = v
	return p
}

func NewPixel(r, g, b, a uint8) Pixel {
	return Pixel{NRGBA: color.NRGBA{r, g, b, a}}
}

type Image struct {
	indexRef int // for subimage
	stride   int // for subimage
	width    int
	height   int
	pixels   []Pixel
}

func NewImage(w, h int) Image {
	return Image{
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

func (i *Image) SubImage(x, y, w, h int) Image {
	r := Image{
		pixels:   i.pixels,
		width:    w,
		height:   h,
		stride:   i.stride,
		indexRef: i.indexRef + y*i.stride + x,
	}
	return r
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.NRGBAModel
}

func (i Image) At(x, y int) color.Color {
	index := i.indexRef + y*i.stride + x
	return i.pixels[index]
}
