package valid_parentheses

var (
	set = map[rune]struct{}{
		'(': {},
		'[': {},
		'{': {},
		'}': {},
		']': {},
		')': {},
	}
	opens = map[rune]struct{}{
		'(': {},
		'[': {},
		'{': {},
	}
	closes = map[rune]struct{}{
		')': {},
		']': {},
		'}': {},
	}
	pairs = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
)

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	var (
		pushed rune
		count  int
	)
	for _, ch := range s {
		if _, ok := set[ch]; !ok {
			return false
		}
		if _, ok := opens[ch]; ok {
			stack = append(stack, ch)
			count++
			continue
		}
		if _, ok := closes[ch]; !ok || count == 0 {
			return false
		}
		if pushed = stack[count-1]; pushed != pairs[ch] {
			return false
		}
		stack = stack[:count-1]
		count--
	}
	return count == 0
}
