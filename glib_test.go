package glib

import (
	"fmt"
	"image"
	"testing"
	// "github.com/labstack/gommon/color"
)

func TestImageImplementImage(t *testing.T) {
	var _ image.Image = NewImage(10, 10)
}

func TestSubPixel(t *testing.T) {
	// Image 3x4
	// . . . .
	// . # # #
	// . # # #
	img := NewImage(4, 3)

	red := RED
	blue := BLUE

	// first row
	img.Set(0, 0, red)
	img.Set(1, 0, red)
	img.Set(2, 0, red)
	img.Set(3, 0, red)
	// second row
	img.Set(0, 1, red)
	img.Set(1, 1, blue)
	img.Set(2, 1, blue)
	img.Set(3, 1, blue)
	// third row
	img.Set(0, 2, red)
	img.Set(1, 2, blue)
	img.Set(2, 2, blue)
	img.Set(3, 2, blue)

	subImage := img.SubImage(1, 1, 3, 3)
	pixels := subImage.SubPixels()
	fmt.Println("len:", len(pixels))
	for index := 0; index < len(pixels); index += 4 {
		fmt.Println("i:", index/4, pixels[index], pixels[index+1], pixels[index+2], pixels[index+3])
	}

}
