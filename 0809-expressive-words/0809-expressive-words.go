func expressiveWords(s string, words []string) int {
	if len(s) == 0 || len(words) == 0 {
		return 0
	}
	var count int
	for _, word := range words {
		if stretchy(s, word) {
			count++
		}
	}
	return count
}

func stretchy(s, word string) bool {
	var i, j int
	for i < len(s) && j < len(word) {
		if s[i] != word[j] {
			return false
		}
		len1, len2 := getRepeatedLength(s, i), getRepeatedLength(word, j)
		if (len1 < 3 && len1 != len2) || (len1 >= 3 && len1 < len2) {
			return false
		}
		i, j = i+len1, j+len2
	}
	return i == len(s) && j == len(word)
}

func getRepeatedLength(str string, i int) int {
	j := i + 1
	for ; j < len(str) && str[i] == str[j]; j++ {}
	return j - i
}