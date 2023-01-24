package glib

import (
	"image"
	"testing"
)

func TestImageImplementImage(t *testing.T) {
	var _ image.Image = NewImage(10, 10)
}
