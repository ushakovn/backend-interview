package group_anagrams

import "sort"

func groupAnagramsV1(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		runes := []rune(str)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sorted := string(runes)

		if _, ok := groups[sorted]; !ok {
			groups[sorted] = []string{str}
			continue
		}

		groups[sorted] = append(groups[sorted], str)
	}
	out := make([][]string, 0, len(groups))

	for _, group := range groups {
		out = append(out, group)
	}
	return out
}

func groupAnagramsV2(strs []string) [][]string {
	var (
		sets   []map[rune]struct{}
		groups [][]string
	)

	for _, str := range strs {
		cur := make(map[rune]struct{})

		for _, ch := range str {
			cur[ch] = struct{}{}
		}

		var (
			ok  bool
			idx int
		)

		for _, set := range sets {

			if len(set) != len(cur) {
				continue
			}

			for ch := range cur {
				if _, ok = set[ch]; !ok {
					idx++
					break
				}
			}

			if ok {
				break
			}
		}

		if ok {
			groups[idx] = append(groups[idx], str)

		} else {
			groups = append(groups, []string{str})

			sets = append(sets, cur)
		}
	}

	return groups
}
