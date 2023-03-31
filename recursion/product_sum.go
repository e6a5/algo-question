package recursion

/*
array = [5, 2, [7, -1], 3, [6, [-13, 8], 4]
output = 12 calculated as 5 + 2 + 2 *(7 -1) + 3 + 2 * ( 6 + 3 * (-13 + 8) + 4)
*/

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	// Write your code here.
	return CountProductValue(array, 1)
}

func CountProductValue(array SpecialArray, depth int) int {
	sum := 0
	for _, product := range array {
		switch v := product.(type) {
		case int:
			sum += v
		case SpecialArray:
			sum += CountProductValue(v, depth+1)
		}
	}
	return depth * sum
}
