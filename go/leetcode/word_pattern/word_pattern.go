package word_pattern

func wordPattern(pattern string, s string) bool {
	var (
		words []string
		word  []rune
	)
	for index, ch := range s {
		if ch == ' ' && len(word) > 0 {
			words = append(words, string(word))
			word = nil
		} else {
			word = append(word, ch)
		}
		if index == len([]rune(s))-1 && len(word) > 0 {
			words = append(words, string(word))
			word = nil
		}
	}
	if len([]rune(pattern)) != len(words) {
		return false
	}
	ruleForward := map[rune]string{}
	ruleBackward := map[string]rune{}

	for index, pt := range pattern {
		word := words[index]

		if _, ok := ruleForward[pt]; !ok {
			ruleForward[pt] = word
		}
		if _, ok := ruleBackward[word]; !ok {
			ruleBackward[word] = pt
		}
		if ruleForward[pt] != word || ruleBackward[word] != pt {
			return false
		}
	}
	return true
}
