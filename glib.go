package glib

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

type Image struct {
	indexRef int // for subimage
	stride   int // for subimage
	width    int
	height   int
	pixels   []byte
}

func (i *Image) Pixels() []byte {
	return i.pixels
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
		pixels: make([]byte, 4*w*h, 4*w*h),
		width:  w,
		height: h,
		stride: w,
	}
}

//go:inline
func (i *Image) GetIndex(x, y int) int {
	index := i.indexRef + y*i.stride*4 + x*4
	return index

}

func (i *Image) DrawHLine(x0, x1, y int, c color.Color) {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	for x := x0; x <= x1; x++ {
		i.Set(x, y, c)
	}
}

func (i *Image) DrawVLine(y0, y1, x int, c color.Color) {
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	for y := y0; y <= y1; y++ {
		i.Set(x, y, c)
	}
}

func (i *Image) Set(x, y int, c color.Color) {
	index := i.GetIndex(x, y)

	// we store color as non-alpha-premultiplied
	switch v := c.(type) {
	case color.RGBA:
		i.pixels[index] = v.R
		i.pixels[index+1] = v.G
		i.pixels[index+2] = v.B
		i.pixels[index+3] = v.A
	default:
		res := color.RGBAModel.Convert(c)
		res2 := res.(color.RGBA)
		i.pixels[index] = res2.R
		i.pixels[index+1] = res2.G
		i.pixels[index+2] = res2.B
		i.pixels[index+3] = res2.A

	}
}

func (i *Image) Fill(c color.Color) *Image {
	for x := 0; x < i.width; x++ {
		for y := 0; y < i.height; y++ {
			i.Set(x, y, c)
		}
	}
	return i
}

func (i *Image) SubImageFromRect(rect image.Rectangle) *Image {
	w := rect.Dx()
	h := rect.Dy()
	return i.SubImage(rect.Min.X, rect.Min.Y, w, h)
}

func (i *Image) SubImage(x, y, w, h int) *Image {
	r := &Image{
		pixels:   i.pixels,
		width:    w,
		height:   h,
		stride:   i.stride,
		indexRef: i.GetIndex(x, y),
	}
	return r
}

// implement image.Image

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	index := i.GetIndex(x, y)
	return color.RGBA{
		R: i.pixels[index],
		G: i.pixels[index+1],
		B: i.pixels[index+2],
		A: i.pixels[index+3],
	}
}

func (i *Image) ToPng(w io.Writer) error {
	return png.Encode(w, i)
}

func (i *Image) ToPngFile(path string) error {
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	err = i.ToPng(fd)
	if err != nil {
		fd.Close()
		return err
	}
	fd.Close()
	return nil
}
