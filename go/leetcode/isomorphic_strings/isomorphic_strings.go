package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	count := len(s)

	if count != len(t) {
		return false
	}
	sr := []rune(s)
	tr := []rune(t)

	rulesST := make(map[rune]rune, count)
	rulesTS := make(map[rune]rune, count)

	for idx := 0; idx < count; idx++ {
		ch := sr[idx]

		if _, ok := rulesST[ch]; !ok {
			rulesST[ch] = tr[idx]
		}
		ch = tr[idx]

		if _, ok := rulesTS[ch]; !ok {
			rulesTS[ch] = sr[idx]
		}
	}
	for idx := 0; idx < count; idx++ {
		if rulesST[sr[idx]] != tr[idx] || rulesTS[tr[idx]] != sr[idx] {
			return false
		}
	}
	return true
}
