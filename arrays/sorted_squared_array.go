package main

import (
	"math"
)

/*
input = [-5, -3, 1, 2, 4]
output = [1, 4, 9, 16, 25]
*/

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	i := 0
	j := len(array) - 1
	idx := len(array) - 1
	newArray := make([]int, len(array))
	for i <= j {
		largerIdx := j
		larger := array[j]
		if math.Abs(float64(array[i])) > math.Abs(float64(array[j])) {
			larger = array[i]
			largerIdx = i
		}
		newArray[idx] = larger * larger
		idx--
		if largerIdx < j {
			i++
			continue
		}
		j--
	}
	return newArray
}
