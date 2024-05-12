package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}
	for _, ch := range magazine {
		m[ch]++
	}
	var (
		count int
		ok    bool
	)
	for _, ch := range ransomNote {
		if count, ok = m[ch]; !ok || count == 0 {
			return false
		}
		m[ch]--
	}
	return true
}

func test(a, b, _ string) {

}
