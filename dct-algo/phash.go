package main

import (
	"os"
	"github.com/nfnt/resize"
	"image/png"
)

func obtainDCT(filename string) [][]float64 {
	file, _ := os.Open(filename)
	img, _ := png.Decode(file)
	defer file.Close()

	resized := resize.Resize(32, 32, img, resize.Lanczos3)
	imgArr := make([][]float64, 32)
	for x := 0; x < 32; x ++ {
		for y := 0; y < 32; y ++ {
			pixel := resized.At(x, y)
			r, g, b, _ := pixel.RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			imgArr[x] = append(imgArr[x], float64((lum / 256) - 128))
		}
	}
	dctArr := dctEquation(imgArr)
	return dctArr
}

func phash(dct [][]float64) string {
	var runningSum float64
	runningSum = 0
	for x := 0; x < 8; x ++ {
		for y := 0; y < 8; y ++ {
			if x == 0 && y == 0 {
				continue
			} else {
				runningSum += dct[x][y]
			}
		}
	}
	avgDCT := runningSum / 63
	output := ""
	for x := 0; x < len(dct); x ++ {
		for y := 0; y < len(dct[0]); y ++ {
			if dct[x][y] < avgDCT {
				output = output + "0"
			} else {
				output = output + "1"
			}
		}
	}
	return output
}

func hammingDistance(val1 string, val2 string) int {
	distance := 0
	for x := 0; x < len(val1); x ++ {
		if (val1[x] != val2[x]) {
			distance ++
		}
	}
	return distance
}