package glib

import (
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPng(t *testing.T) {
	fd, err := os.Create("file.png")
	assert.NoError(t, err)
	defer fd.Close()

	w := 10000
	h := 10000

	img := NewImage(w, h)
	img.Fill(NewPixel(255, 0, 0, 255))
	sub := img.SubImage(25, 25, w-50, h-50)
	sub.Fill(NewPixel(0, 255, 0, 255))
	err = png.Encode(fd, img)
	assert.NoError(t, err)
}
