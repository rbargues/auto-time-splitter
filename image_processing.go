package main

import (
	"os"
	"fmt"
	"image/png"
	"image/color"
)
type PixelCount struct {
	R map[uint8]int
	G map[uint8]int
	B map[uint8]int
}

func main() {
	imageFile, err := os.Open("./Beginning.png")
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	imageInfo, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	bounds := imageInfo.Bounds()
	imageCount := PixelCount{
		R: make(map[uint8]int),
		G: make(map[uint8]int),
		B: make(map[uint8]int),
	}
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba:= imageInfo.At(x,y).(color.RGBA)
			if _, ok := imageCount.R[rgba.R]; ok {
				imageCount.R[rgba.R] += 1
			} else {
				imageCount.R[rgba.R] = 1
			}
			if _, ok := imageCount.G[rgba.G]; ok {
				imageCount.G[rgba.G] += 1
			} else {
				imageCount.G[rgba.G] = 1
			}
			if _, ok := imageCount.B[rgba.B]; ok {
				imageCount.B[rgba.B] += 1
			} else {
				imageCount.B[rgba.B] = 1
			}
		}
	}
	if imageCount.R[255] > 1000 && imageCount.R[255] > 1000 && imageCount.R[255] > 1000 {
		fmt.Println("probably load screen")
	} else {
		fmt.Println("probably other")
	}
	// fmt.Println(imageCount.R[255])
}