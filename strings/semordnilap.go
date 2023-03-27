package strings

// words = ["diaper", "abc", "test", "cba", "repaid"]
// output = [["diaper", "repaid"], ["abc", "cba"]]

func Semordnilap(words []string) [][]string {
	// Write your code here.
	semordnilapList := make(map[string]int)
	result := make([][]string, 0)
	for currentIdx, word := range words {
		if idx, found := semordnilapList[word]; found {
			pair := make([]string, 2)
			pair[0], pair[1] = words[idx], word
			result = append(result, pair)
			continue
		}
		reversedWord := reverseString(word)
		semordnilapList[reversedWord] = currentIdx
	}
	return result
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
