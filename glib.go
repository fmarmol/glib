package glib

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"unsafe"

	"github.com/disintegration/imaging"
)

type Image struct {
	indexRef int             // for subimage
	stride   int             // for subimage
	rect     image.Rectangle // rectangle of the image
	pixels   []byte
}

func (i *Image) Width() int {
	return i.rect.Dx()
}

func (i *Image) Heigth() int {
	return i.rect.Dy()
}

func (i *Image) SubPixels() []byte {
	w := i.rect.Dx()
	h := i.rect.Dy()
	// log.Println("w:", w, "h:", h, w*h)
	ret := make([]byte, 0, 4*i.rect.Dx()*i.rect.Dy())
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// log.Println("x:", x, "y:", y)
			idx := i.GetIndex(x, y)
			ret = append(ret, i.pixels[idx])
			ret = append(ret, i.pixels[idx+1])
			ret = append(ret, i.pixels[idx+2])
			ret = append(ret, i.pixels[idx+3])
		}
	}
	return ret
}

func (i *Image) Pixels() []byte {
	return i.pixels
}

func (i *Image) PixelsU32() []uint32 {
	return *(*[]uint32)(unsafe.Pointer(&i.pixels))
}

func (i *Image) Resize(nw, nh int) *Image {
	w, h := i.rect.Dx(), i.rect.Dy()

	newImg := NewImage(nw, nh)

	for x := 0; x < nw; x++ {
		for y := 0; y < nh; y++ {

			xNorm := (float64(x) + 0.5) / float64(nw)
			yNorm := (float64(y) + 0.5) / float64(nh)

			xSource := int(xNorm * float64(w))
			ySource := int(yNorm * float64(h))
			cSource := i.At(xSource, ySource)

			newImg.Set(x, y, cSource)

		}
	}
	return newImg
}

func (i *Image) Scale(v float64) *Image {
	w, h := i.rect.Dx(), i.rect.Dy()
	ret := NewImage(int(float64(w)*v), int(float64(h)*v))
	w2, h2 := ret.rect.Dx(), ret.rect.Dy()

	// TODO add 0.5 to be in the center of the pixel
	for x := 0; x < ret.rect.Dx(); x++ {
		for y := 0; y < ret.rect.Dy(); y++ {
			xi := int(float64(x) / float64(w2) * float64(w))
			yi := int(float64(y) / float64(h2) * float64(h))
			c := i.At(xi, yi)
			ret.Set(x, y, c)
		}
	}
	return ret
}

func (i *Image) Translate(dx, dy int) *Image {
	i.rect.Min.X += dx
	i.rect.Max.X += dx
	i.rect.Min.Y += dy
	i.rect.Max.Y += dy
	return i
}

func (i *Image) Center() (int, int) {
	return (i.rect.Max.X + i.rect.Min.X) / 2, (i.rect.Max.Y + i.rect.Min.Y) / 2
}

func NewImageFromPath(filename string) *Image {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return NewImageFromPngBytes(data)
}

func NewImageFromPngBytes(data []byte) *Image {
	buf := bytes.NewBuffer(data)
	res, err := png.Decode(buf)
	if err != nil {
		panic(err)
	}
	return NewImageFromImage(res)
}

func NewImageFromBytes(w, h int, pixels []byte) *Image {
	return &Image{
		stride: w,                      // for subimage
		rect:   image.Rect(0, 0, w, h), // rectangle of the image
		pixels: pixels,
	}

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
		rect:   image.Rect(0, 0, w, h),
		stride: w,
	}
}

//go:inline
func (i *Image) GetIndex(x, y int) int {
	index := i.indexRef + y*i.stride*4 + x*4
	// log.Println("HERE3:", index)
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
	case color.NRGBA:
		i.pixels[index] = v.R
		i.pixels[index+1] = v.G
		i.pixels[index+2] = v.B
		i.pixels[index+3] = v.A
	default:
		res := color.NRGBAModel.Convert(c)
		res2 := res.(color.NRGBA)
		i.pixels[index] = res2.R
		i.pixels[index+1] = res2.G
		i.pixels[index+2] = res2.B
		i.pixels[index+3] = res2.A

	}
}

func (i *Image) Fill(c color.Color) *Image {
	w := i.rect.Dx()
	h := i.rect.Dy()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
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

	dx := x + w
	if x+w > i.rect.Max.X {
		dx = i.rect.Max.X
	}
	dy := y + h
	if y+h > i.rect.Max.Y {
		dy = i.rect.Max.Y
	}

	r := &Image{
		pixels:   i.pixels,
		rect:     image.Rect(x, y, dx, dy),
		stride:   i.stride,
		indexRef: i.GetIndex(x, y),
	}
	return r
}

// implement image.Image

func (i *Image) Bounds() image.Rectangle {
	return i.rect
}

func (i *Image) ColorModel() color.Model {
	return color.NRGBAModel
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

func (i *Image) Rotate(angle float64, bgColor color.Color) {
	res := imaging.Rotate(i, angle, bgColor)
	*i = *NewImageFromImage(res)
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
