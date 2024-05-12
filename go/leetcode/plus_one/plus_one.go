package plus_one

func plusOne(digits []int) []int {
	for idx, next := len(digits)-1, true; idx >= 0 && next; idx-- {
		if digits[idx] == 9 {

			digits[idx] = 0

			if idx == 0 {
				idx++

				buf := make([]int, 1)
				digits = append(buf, digits...)
			}
		} else {
			digits[idx]++

			next = false
		}
	}
	return digits
}
