package glib

import "image/color"

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

func NewPixel(r, g, b, a uint8) *Pixel {
	return &Pixel{NRGBA: color.NRGBA{r, g, b, a}}
}
