package main

import (
	"fmt"
)

func main() {
	dctArr := obtainDCT("bobomb.1.png")
	seconddctArr := obtainDCT("chain.1.png")
	// fmt.Printf("%v", dctArr)
	phashVal := phash(dctArr)
	secondphashVal := phash(seconddctArr)


	hammingDist := hammingDistance(phashVal, secondphashVal)
	fmt.Printf("%v", hammingDist)
	fmt.Printf("\n")

	// fmt.Printf("%v", hammingDistance(phashVal, secondphashVal))
}