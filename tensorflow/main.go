package main

import (
	"fmt"
	"image"
	"image/png"
	"image/color"
	"os"
	"github.com/kbinani/screenshot"
	"time"
)

func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func screenshotGrab(filepath string) {
	img, err := screenshot.Capture(95, 237, 850, 580)
	if err != nil {
		panic(err)
	}
	save(img, filepath)
}
// type PixelCount struct {
// 	R map[uint8]int
// 	G map[uint8]int
// 	B map[uint8]int
// }
func imageProcess() bool{
	imageFile, err := os.Open("./screenshot.png")
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	imageInfo, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	bounds := imageInfo.Bounds()
	// imageCount := PixelCount{
	// 	R: make(map[uint8]int),
	// 	G: make(map[uint8]int),
	// 	B: make(map[uint8]int),
	// }
	blackPixels := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba:= imageInfo.At(x,y).(color.RGBA)
			if rgba.R < 10 && rgba.G < 10 && rgba.B < 10 {
				blackPixels ++
			}
			// if _, ok := imageCount.R[rgba.R]; ok {
			// 	imageCount.R[rgba.R] += 1
			// } else {
			// 	imageCount.R[rgba.R] = 1
			// }
			// if _, ok := imageCount.G[rgba.G]; ok {
			// 	imageCount.G[rgba.G] += 1
			// } else {
			// 	imageCount.G[rgba.G] = 1
			// }
			// if _, ok := imageCount.B[rgba.B]; ok {
			// 	imageCount.B[rgba.B] += 1
			// } else {
			// 	imageCount.B[rgba.B] = 1
			// }
		}
	}
	pixelsTotal := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
	percent := float64(blackPixels) / float64(pixelsTotal)
	if percent > 0.95 {
		return true
	}
	return false
}
func main() {
	screenshotCount := 359
	for true {
		time.Sleep(250 * time.Millisecond)
		screenshotGrab("screenshot.png")
		blackScreen := imageProcess()
		if blackScreen == true {
			time.Sleep(1 * time.Second)
			endTime := time.Now().Add(2 * time.Second)
			for time.Now().Before(endTime) {
				fileName := fmt.Sprintf("%v.png", screenshotCount)
				screenshotCount ++
				screenshotGrab(fileName)
				time.Sleep(250 * time.Millisecond)
			}
			fmt.Println("got a loading screen")
		}
	}
}