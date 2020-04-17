package main

import (
	"image"
	"image/png"
	"image/color"
	"image/draw"
	"os"
	// "fmt"
)
/*
for world name on 630, 470 img
imgWidth 230
imgHeight 50
point 200,380

for star count 630, 470 img
imgWidth 30
imgHeight 50
point 390,20

for star name 630, 470 img
imgWidth 300
imgHeight 50
point 140, 150

for couse number 630, 470 img
imgWidth 100
imgHeight 40
point 120, 110
*/
func main() {
	
	// img := image.NewRGBA(image.Rectangle{image.Point{0,0}, image.Point{imgWidth, imgHeight}})
	// imgColor := color.RGBA{255,255,255,255}
	// for x := 0; x < imgWidth; x ++ {
	// 	for y := 0; y < imgHeight; y ++ {
	// 		img.Set(x, y, imgColor)
	// 	}
	// }
	// file, _ := os.Create("canvas.png")
	// png.Encode(file, img)
	// defer file.Close()

	// canvas, _ := os.Open("./canvas.png")
	// defer canvas.Close()
	// canvasInfo, _ := png.Decode(canvas)

	cropCopy, _ := os.Open("./obtainedstar.png")
	//630, 470
	defer cropCopy.Close()
	cropInfo, _ := png.Decode(cropCopy)
	imgWidth := 100
	imgHeight := 40

	m := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(m, image.Rect(0, 0, imgWidth, imgHeight), cropInfo, image.Point{120,110}, draw.Src)
	imgColor := color.RGBA{0,0,0,255}
	bounds := m.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba := m.At(x,y).(color.RGBA)
			if (rgba.R < 64 && rgba.R > 55) && (rgba.G < 64 && rgba.G > 55) && (rgba.B < 64 && rgba.B > 55) {
				continue;
			}
			if rgba.R	< 180 && rgba.G < 180 && rgba.B < 180 {
				m.Set(x, y, imgColor)
			} else if rgba.B < 150 {
				m.Set(x, y, imgColor)
			}
		}
	}


	newImg, _ := os.Create("copied.png")
	defer newImg.Close()

	png.Encode(newImg, m)
	// func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	// r := image.Rectangle{dp, dp.Add(sr.Size())}
	// draw.Draw(dst, r, src, sr.Min, draw.Src)
	// crop := image.Rectangle{image.Point{0,0}, }
}