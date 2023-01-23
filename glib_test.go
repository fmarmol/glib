package glib

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPixel(t *testing.T) {
	p := NewPixel(1, 2, 3, 4)
	// test R
	assert.Equal(t, p.R(), uint8(1))
	p.SetR(41)
	assert.Equal(t, p.R(), uint8(41))
	// test G
	assert.Equal(t, p.G(), uint8(2))
	p.SetG(42)
	assert.Equal(t, p.G(), uint8(42))
	// test B
	assert.Equal(t, p.B(), uint8(3))
	p.SetB(43)
	assert.Equal(t, p.B(), uint8(43))
	// test A
	assert.Equal(t, p.A(), uint8(4))
	p.SetA(44)
	assert.Equal(t, p.A(), uint8(44))
}

func TestPixelImplementColor(t *testing.T) {
	var _ color.Color = NewPixel(1, 2, 3, 4)
}

func TestImageImplementImage(t *testing.T) {
	var _ image.Image = NewImage(10, 10)
}
