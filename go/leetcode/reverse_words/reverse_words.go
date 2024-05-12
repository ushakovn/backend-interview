package reverse_words

import "unicode"

func reverseWords(s string) string {
	reverse := make([]rune, 0, len(s))
	stack := make([]rune, 0)

	flush := func() {
		for count := len(stack); count != 0; count = len(stack) {
			reverse = append(reverse, stack[count-1])
			stack = stack[:count-1]
		}
	}
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			stack = append(stack, ch)
			continue
		}
		flush()

		reverse = append(reverse, ch)
	}
	flush()

	return string(reverse)
}
