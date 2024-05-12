package valid_anagram

func isAnagram(s string, t string) bool {
	sr := []rune(s)
	st := []rune(t)

	count := len(sr)

	if count != len(st) {
		return false
	}
	ms := map[rune]int{}
	mt := map[rune]int{}

	m := map[rune]struct{}{}

	var (
		chs []rune
		ch  rune
	)
	for idx := 0; idx < count; idx++ {
		ch = sr[idx]
		ms[ch]++

		if _, ok := m[ch]; !ok {
			chs = append(chs, ch)
			m[ch] = struct{}{}
		}
		ch = st[idx]
		mt[ch]++

		if _, ok := m[ch]; !ok {
			chs = append(chs, ch)
			m[ch] = struct{}{}
		}
	}
	for _, ch := range chs {
		if ms[ch] != mt[ch] {
			return false
		}
	}
	return true
}
