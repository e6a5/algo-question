package greedyalgorithms

import (
	"sort"
)

// queries = [3, 2, 1, 2, 6]
// output = 17

func MinimumWaitingTime(queries []int) int {
	// Write your code here.
	sort.Slice(queries, func(a, b int) bool {
		return queries[b] < queries[a]
	})

	sum := 0
	for i, num := range queries {
		sum += i * num
	}

	return sum
}
