package main

import (
	"fmt"
	"strings"
)

func main() {
	var first_string string = "david"
	first_string_array := strings.Split(strings.ToLower(first_string),"")
	var second_string string = "mario"
	second_string_array := strings.Split(strings.ToLower(second_string),"")
	matrix := make([][]int, len(first_string) + 1)
	for i := 0; i <= len(first_string); i++ {
		matrix[i] = make([]int, len(second_string) + 1)
	}
	for i := 0; i <= len(first_string); i++ {
		matrix[i][0] = i
	}
	for i := 0; i <= len(second_string); i++ {
		matrix[0][i] = i
	}
	for i := 1; i <= len(first_string); i++ {
		for j := 1; j <= len(second_string); j++ {
			var subCost int
			if first_string_array[i - 1] == second_string_array[j - 1] {
				subCost = 0	
				_ = subCost
			} else {
				subCost = 1
			}
			deletion := matrix[i - 1][j] + 1
			insertion := matrix[i][j - 1] + 1
			substitution := matrix[i - 1][j - 1] + subCost
			if deletion <= insertion && deletion <= substitution {
				matrix[i][j] = deletion
			} else if insertion <= deletion && insertion <= substitution {
				matrix[i][j] = insertion
			} else {
				matrix[i][j] = substitution
			}
		}
	}
	fmt.Printf("%v\n", matrix)
}