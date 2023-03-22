package main

func InsertionSort(array []int) []int {
	// Write your code here.
	n := len(array)
	for i := 1; i < n; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && key < array[j] {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key

	}
	return array
}
