package length_of_last_word

import "unicode"

func lengthOfLastWord(s string) int {
	var (
		last  int
		count int
	)
	for _, ch := range s {
		if unicode.IsSpace(ch) {
			if count != 0 {
				last = count
				count = 0
			}
			continue
		}
		count++
	}
	if count == 0 {
		return last
	}
	return count
}
