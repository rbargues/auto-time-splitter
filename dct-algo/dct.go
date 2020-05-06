package main

import (
	"math"
)
//References:
//https://www.math.cuhk.edu.hk/~lmlui/dct.pdf
/*
taking in an array of arrays block shape 8x8 or 32x32 ex.
[
	[],
	[],
	[],
]
*/
func dctEquation(input [][]float64) [][]float64 {
	output := make([][]float64, len(input))
	// in an 8x8 block represents 1 / sqrt(2 * N) [n = 8]
	constant := 1 / math.Sqrt(2 * float64(len(input)))
	for row := 0; row < len(input); row ++ {
		for col := 0; col < len(input[0]); col ++ {
			var c1 float64
			var c2 float64
			if row == 0 {
				c1 = 1 / math.Sqrt(2)
			} else {
				c1 = 1
			}
			if col == 0 {
				c2 = 1 / math.Sqrt(2)
			} else {
				c2 = 1
			}

			value := input[row][col]

			var runningSum float64
			runningSum = 0
			for x := 0; x < len(input); x++ {
				for y := 0; y < len(input[0]); y++ {
					var xCos float64
					var yCos float64
					xCos = (((2 * float64(x)) + 1) * (float64(row) * math.Pi)) / (2 * float64(len(input)))
					yCos = (((2 * float64(y)) + 1) * (float64(col) * math.Pi)) / (2 * float64(len(input)))

					runningSum += value * math.Cos(xCos) * math.Cos(yCos)
				}
			}
			runningSum = runningSum * constant * c1 * c2
			output[row] = append(output[row], runningSum)
		}
	}
	return output 
}
