package is_subsequence

func isSubsequence(s string, t string) bool {
	sr := []rune(s)
	tr := []rune(t)

	lsr := len(sr)
	ltr := len(tr)

	if lsr == 0 {
		return true
	}
	if lsr > ltr {
		return false
	}
	var jdx int

	for idx := 0; idx < ltr; idx++ {
		if tr[idx] == sr[jdx] {
			jdx++
		}
		if jdx == lsr {
			return true
		}
	}
	return false
}
