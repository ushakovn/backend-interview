package first_occurrence_string

func strStr(haystack string, needle string) int {
	haystackCh := []rune(haystack)
	needleCh := []rune(needle)
	var (
		jdx int
		out = -1
	)
	haystackCount := len(haystackCh)
	needleCount := len(needleCh)

	for idx := 0; needleCount <= haystackCount && idx < haystackCount && jdx < needleCount; idx++ {
		if haystackCh[idx] != needleCh[jdx] {
			if out != -1 {
				idx = out
			}
			jdx = 0
			out = -1
			continue
		}
		if out == -1 {
			out = idx
		}
		jdx++
	}
	if jdx != needleCount {
		return -1
	}
	return out
}
