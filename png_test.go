package glib

import (
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPng(t *testing.T) {
	fd, err := os.Create("file.png")
	assert.NoError(t, err)
	defer fd.Close()

	w := 1000
	h := 1000

	img := NewImage(w, h)
	img.Fill(color.RGBA{255, 0, 0, 255})
	sub := img.SubImage(25, 25, w-50, h-50)
	sub.Fill(color.RGBA{0, 255, 0, 255})
	img.DrawHLine(300, 600, 400, color.Black)
	img.DrawVLine(300, 600, 400, color.White)
	err = img.ToPng(fd)
	assert.NoError(t, err)

	fd.Close()
	fd, err = os.Open("file.png")
	assert.NoError(t, err)
	imgCopy, err := png.Decode(fd)
	assert.NoError(t, err)

	img2 := NewImageFromImage(imgCopy)
	fd2, err := os.Create("file2.png")
	assert.NoError(t, err)
	defer fd2.Close()
	err = img2.ToPng(fd2)
	assert.NoError(t, err)

	data, err := ioutil.ReadFile("file.png")
	assert.NoError(t, err)
	data2, err := ioutil.ReadFile("file2.png")
	assert.NoError(t, err)
	assert.Equal(t, data, data2)
}
