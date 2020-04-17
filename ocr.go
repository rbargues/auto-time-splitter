//ocr example from https://github.com/otiai10/gosseract
package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("./bobomb.png")
	text, _ := client.Text()
	output_text := fmt.Sprintln(text)
	fmt.Printf("%v",output_text)
}