package palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	runes := []rune(strconv.Itoa(x))
	count := len(runes)

	for idx := 0; idx < count; idx++ {
		if runes[idx] != runes[count-idx-1] {
			return false
		}
	}
	return true
}
