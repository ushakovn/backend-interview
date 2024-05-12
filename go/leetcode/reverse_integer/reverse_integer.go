package reverse_integer

func reverse(x int) int {
	var (
		buf      []int
		digits   []int
		negative bool
	)
	if x < 0 {
		x *= -1
		negative = true
	}
	count := 1

	for x != 0 {
		buf = append(buf, x%10)
		digits = append(digits, count)
		count *= 10
		x /= 10
	}
	out := 0

	for idx := len(buf); idx > 0; idx-- {
		out += buf[idx-1] * digits[len(digits)-idx]

		if !validInt32(out) {
			return 0
		}
	}
	if negative {
		out *= -1
	}
	return out
}

func validInt32(x int) bool {
	const maxInt32 = 1<<31 - 1
	return x <= maxInt32
}
