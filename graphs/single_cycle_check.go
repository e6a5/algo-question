package graphs

/**
* a single cycle occurs if, starting at any index in the array and following the jumps,
* very element in the array is visted exactly once before landing back on the starting index
* input [2,3,1,-4,-4,2]
* output true
 */

func HasSingleCycle(array []int) bool {
	// Write your code here.
	n := len(array)
	startIdx := 1
	currentIdx := startIdx
	vistedMap := make(map[int]bool)
	hasCycled := false
	for {
		jump := array[currentIdx]
		vistedMap[currentIdx] = true
		nextIdx := currentIdx + jump

		if nextIdx < 0 {
			reminder := nextIdx % n
			if reminder == 0 {
				reminder = (-1) * n
			}
			nextIdx = n + reminder
		} else if nextIdx >= n {
			nextIdx = nextIdx % n
		}
		if nextIdx == startIdx {
			break
		}
		if vistedMap[nextIdx] {
			hasCycled = true
			break
		}
		currentIdx = nextIdx
	}
	if hasCycled {
		return false
	}
	if len(vistedMap) == n {
		return true
	}
	return false
}
