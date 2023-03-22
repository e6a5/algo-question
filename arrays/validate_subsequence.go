package main

/*
array = [5, 1, 22, 25, 6, -1, 8, 10]
sequence = [1, 6, -1, 10]
output = true
*/

func IsValidSubsequence(array []int, sequence []int) bool {
	temp := sequence
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == sequence[len(temp)-1] {
			temp = temp[0 : len(temp)-1]
		}
		if len(temp) == 0 {
			return true
		}
	}
	return false
}
