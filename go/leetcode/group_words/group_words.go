package group_words

func groupWords(words []string) [][]string {
	sets := make([]map[rune]struct{}, 0, len(words))
	groups := make([][]string, 0, len(words))

	for _, word := range words {
		cur := map[rune]struct{}{}

		for _, ch := range word {
			cur[ch] = struct{}{}
		}
		var (
			found bool
			index int
		)
		for idx, set := range sets {
			index = idx
			found = true

			for _, ch := range word {
				if _, ok := set[ch]; !ok {
					found = false
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			groups[index] = append(groups[index], word)

			continue
		}
		groups = append(groups, []string{word})
		sets = append(sets, cur)
	}
	return groups
}
