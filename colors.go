package glib

import "image/color"
import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Color = color.NRGBA

func New(i int) Color {
	r, g, b, a := uint8(i>>24&0xFF), uint8(i>>16&0xFF), uint8(i>>8&0xFF), uint8(i&0xFF)
	return Color{r, g, b, a}
}

var (
	RANDOM = func() color.NRGBA {
		return color.NRGBA{
			R: uint8(rand.Int31n(255)),
			G: uint8(rand.Int31n(255)),
			B: uint8(rand.Int31n(255)),
			A: 255,
		}
	}
	GREEN = New(0x00FF00FF)
	RED   = New(0xFF0000FF)
	BLUE  = New(0x0000FFFF)
)
