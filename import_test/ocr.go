package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	// "strconv"
)

func ocr() int {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("copied.png")
	text, _ := client.Text()
	outputText := fmt.Sprintln(text)
	var intVersion int
	_, err := fmt.Sscan(outputText, &intVersion)
	if err != nil {
		
	}
	// fmt.Println("new")
	// fmt.Println(outputText)
	// fmt.Println(intVersion)
	return intVersion
}
