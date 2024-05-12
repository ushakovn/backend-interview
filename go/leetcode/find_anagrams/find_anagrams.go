package find_anagrams

func findAnagrams(s string, p string) []int {
	var indices []int

	chars := map[rune]int{}
	check := map[rune]struct{}{}

	for _, ch := range p {
		chars[ch]++
		check[ch] = struct{}{}
	}
	var (
		buf   map[rune]int
		found bool
		count int
		ind   int
	)
	rs := []rune(s)
	cs := len(rs)

	for i := 0; i < cs; i++ {
		ch := rs[i]

		if _, ok := check[ch]; !ok {
			if found {
				i = ind
				count = 0
				found = false
				buf = map[rune]int{}
			}
			continue
		}
		if !found {
			ind = i
			count = 0
			found = true
			buf = map[rune]int{}
		}
		buf[ch]++
		count++

		if count == len([]rune(p)) {
			for _, ch := range p {
				if chars[ch] != buf[ch] {
					found = false
					break
				}
			}
			if found {
				indices = append(indices, ind)
			}
			i = ind
			count = 0
			found = false
			buf = map[rune]int{}
		}
	}
	return indices
}
