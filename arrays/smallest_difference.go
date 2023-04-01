package arrays

package main

import (
    "sort"
    "math"
)

/*
arrayOne = [-1, 5, 10, 20, 28, 3]
arrayTwo = [26, 134, 135, 15, 17]

sort the input arrays
arrayOne = [-1, 3, 5, 10, 20, 28]
arrayTwo = [15, 17, 26, 134, 135]

smallest = maxfloat64
using two pointers starting at beginning of two arrays

currentDifference := abs(-1 - 15) = 16
currentDifference < smallest => smallest = 16
-1 < 15 => idxOfArrayOne++

currentDifference := abs(3 - 15) = 12
currentDifference < smallest => smallest = 12
3 < 15 => idxOfArrayOne++

currentDifference := abs(5 - 15) = 10
currentDifference < smallest => smallest = 10
5 < 15 => idxOfArrayOne++

currentDifference := abs(10 - 15) = 5
currentDifference < smallest => smallest = 5
10 < 15 => idxOfArrayOne++

currentDifference := abs(20 - 15) = 5
currentDifference = smallest => keep the smallest
20 < 15 => indxOfArrayTwo++

...
28 - 26 = 2 => smallest = 2 => result [28, 26]

*/

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
    sort.Ints(array1)
    sort.Ints(array2)
    oneIdx, twoIdx := 0, 0
    result := []int{}
    smallest := math.MaxFloat64
    for oneIdx < len(array1) && twoIdx< len(array2) {
        current := math.Abs(float64(array1[oneIdx]) - float64(array2[twoIdx]))
        if current < smallest {
            smallest = current
            result = []int{array1[oneIdx], array2[twoIdx]}
        }
        if array1[oneIdx] > array2[twoIdx] {
            twoIdx++
        } else {
            oneIdx++
        }
    }
    return result
}
