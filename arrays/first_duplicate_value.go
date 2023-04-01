package arrays

func FirstDuplicateValue(array []int) int {
	// Write your code here.
	uniqueValue := make(map[int]bool)
	for _, value := range array {
		if uniqueValue[value] {
			return value
		}
		uniqueValue[value] = true
	}
	return -1
}
