package main

import (
	"image"
	"image/png"
	// "image/color"
	"image/draw"
	"os"
)

func imgCrop(){
	cropCopy, _ := os.Open("screenshot.png")
	//630, 470
	defer cropCopy.Close()
	cropInfo, _ := png.Decode(cropCopy)
	imgWidth := 35
	imgHeight := 40

	m := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(m, image.Rect(0, 0, imgWidth, imgHeight), cropInfo, image.Point{540,25}, draw.Src)

	newImg, _ := os.Create("star.png")
	defer newImg.Close()

	png.Encode(newImg, m)
}