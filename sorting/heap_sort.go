package main

func HeapSort(array []int) []int {
	// Write your code here.
	n := len(array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, i, 0)
	}
	return array
}

func heapify(array []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && array[i] < array[l] {
		largest = l
	}
	if r < n && array[largest] < array[r] {
		largest = r
	}
	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, n, largest)
	}
}
