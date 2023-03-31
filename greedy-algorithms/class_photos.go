package greedyalgorithms

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	tracker := redShirtHeights[0] > blueShirtHeights[0]
	for i := 0; i < len(redShirtHeights); i++ {
		if (redShirtHeights[i] >= blueShirtHeights[i]) != tracker {
			return false
		}
	}
	return true
}
