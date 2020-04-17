package main

import (
	"os"
	// "fmt"
	"image"
	"image/draw"
	"image/png"
	"image/color"
)

func processImg(filepath string) {
	imageFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	imageInfo, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	bounds := imageInfo.Bounds()

	m := image.NewRGBA(bounds)
	draw.Draw(m, bounds, imageInfo, image.Point{0,0}, draw.Src)

	imgColor := color.RGBA{0,0,0,255}
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba:= imageInfo.At(x,y).(color.RGBA)
			if rgba.R < 200 || rgba.B > 20 {
				m.Set(x, y, imgColor)
			} 
		}
	}
	newImg, _ := os.Create("copied.png")
	defer newImg.Close()
	png.Encode(newImg, m)
	// fmt.Println(imageCount.R[255])
}