package glib

import (
	"image"
	"image/draw"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDraw(t *testing.T) {

	dst := NewImage(100, 100)

	for x := dst.Bounds().Min.X; x <= dst.Bounds().Max.X-1; x++ {
		for y := dst.Bounds().Min.Y; y <= dst.Bounds().Max.Y-1; y++ {
			dst.Set(x, y, RANDOM())
		}
	}
	src := NewImage(25, 25).Fill(RANDOM())
	src.SubImage(4, 4, 10, 10).Fill(RANDOM())

	draw.Draw(dst, image.Rect(
		10,
		10,
		20,
		20,
	),
		src,
		image.Point{0, 0},
		draw.Over,
	)
	draw.Draw(dst, image.Rect(
		40,
		40,
		50,
		50,
	),
		src,
		image.Point{5, 5},
		draw.Over,
	)

	err := dst.ToPngFile("test.png")
	assert.NoError(t, err)
}
