package reverse_words_2

import "unicode"

func reverseWords(s string) string {
	var (
		words [][]rune
		word  []rune
	)
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			word = append(word, ch)
			continue
		}
		if len(word) != 0 {
			words = append(words, word)
		}
		word = make([]rune, 0)
	}
	if len(word) != 0 {
		words = append(words, word)
	}
	var reversed []rune

	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i]...)
		if i != 0 {
			reversed = append(reversed, ' ')
		}
	}
	return string(reversed)
}
