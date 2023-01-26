package glib

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func (i *Image) RenderText(f font.Face, x, y int, text string, c color.Color) {
	d := font.Drawer{
		Dst:  i,
		Src:  image.NewUniform(c),
		Face: f,
		Dot:  fixed.P(x, y),
	}
	d.DrawString(text)
}
