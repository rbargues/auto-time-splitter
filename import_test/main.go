package main

import (
	"fmt"
)

func main() {
	starCount := 0
	for true {
		screenshotGrab()
		imgCrop()
		processImg("star.png")
		ocrVal := ocr()
		if starCount != ocrVal {
			starCount = ocrVal
			fmt.Println(starCount)
			break
		}
	}
}