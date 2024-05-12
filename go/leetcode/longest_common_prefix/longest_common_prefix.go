package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	var (
		outChars  []rune
		bufChar   rune
		curChar   rune
		charIdx   int
		strFinish bool
	)
	for !strFinish {
		for strIdx, str := range strs {
			if len(str) <= charIdx {
				strFinish = true
				break
			}
			if strIdx == 0 {
				bufChar = rune(str[charIdx])
				continue
			}
			if curChar = rune(str[charIdx]); curChar != bufChar {
				strFinish = true
				break
			}
		}
		if !strFinish && len(outChars) <= charIdx {
			outChars = append(outChars, bufChar)
		}
		charIdx++
	}
	return string(outChars)
}
