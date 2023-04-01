package arrays

/*
array = [3, 5, -4, 8, 11, 1, -1, 6]
targetSum = 10
output = [-1, 11]
*/

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	checkedHash := make(map[int]bool)
	for _, x := range array {
		y := target - x
		_, exists := checkedHash[y]
		if !exists {
			checkedHash[x] = true
			continue
		}
		return []int{x, y}
	}
	return []int{}
}
