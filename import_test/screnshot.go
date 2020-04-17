package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"github.com/kbinani/screenshot"
)

func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func screenshotGrab() {
	img, err := screenshot.Capture(72, 59, 630, 470)
	if err != nil {
		panic(err)
	}
	filepath := fmt.Sprintf("screenshot.png")
	save(img, filepath)
}