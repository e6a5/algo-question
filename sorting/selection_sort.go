package main

func SelectionSort(array []int) []int {
	// Write your code here.
	n := len(array)
	for i := 0; i < n-1; i++ {
		minimum := array[i]
		minimumIdx := i
		for j := i + 1; j < n; j++ {
			if array[j] < minimum {
				minimum = array[j]
				minimumIdx = j
			}
		}
		if minimumIdx != i {
			array[i], array[minimumIdx] = array[minimumIdx], array[i]
		}
	}
	return array
}
