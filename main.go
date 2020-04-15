package main

import (
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func main() {
	// Capture each displays.
	bound := screenshot.GetDisplayBounds(0)
	img, err := screenshot.Capture(bound.Min.X, bound.Min.Y, bound.Dx(), bound.Dy())
	if err != nil {
		panic(err)
	}
	save(img, "all.png")
}
