package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
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
	bound := screenshot.GetDisplayBounds(0)
	count := 1
	for count < 10 {
		time.Sleep(1 * time.Second)
		count ++
		img, err := screenshot.Capture(bound.Min.X, bound.Min.Y, bound.Dx(), bound.Dy())
		if err != nil {
			panic(err)
		}
		filepath := fmt.Sprintf("all%v.png", count)
		save(img, filepath)
	}
	
}
