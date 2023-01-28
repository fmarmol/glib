package glib

import "image/color"
import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	RANDOM = func() color.RGBA {
		return color.RGBA{
			R: uint8(rand.Int31n(255)),
			G: uint8(rand.Int31n(255)),
			B: uint8(rand.Int31n(255)),
			A: 255,
		}
	}
)
