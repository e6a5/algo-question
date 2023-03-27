package strings

/*
input = "abcacdde"
output = 1
*/

func FirstNonRepeatingCharacter(str string) int {
	// Write your code here.
	countChars := make(map[rune]int)
	for _, char := range str {
		countChars[char] = countChars[char] + 1
	}
	for i := 0; i < len(str); i++ {
		if countChars[rune(str[i])] == 1 {
			return i
		}
	}
	return -1
}
