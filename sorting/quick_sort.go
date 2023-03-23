package main

func QuickSort(array []int) []int {
	// Write your code here.
	quickSortHelper(array, 0, len(array)-1)
	return array
}

func quickSortHelper(array []int, low int, high int) {
	if low < high {
		pi := partition(array, low, high)
		quickSortHelper(array, low, pi-1)
		quickSortHelper(array, pi+1, high)
	}
}

func partition(array []int, low int, high int) int {
	pivot := array[high]
	i := low - 1
	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			swap(array, i, j)
		}
	}
	swap(array, i+1, high)
	return i + 1
}

func swap(array []int, a, b int) {
	array[a], array[b] = array[b], array[a]
}
