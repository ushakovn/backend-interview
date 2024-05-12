package add_digits

func addDigits(num int) int {
	buf := 0
	n := 0

	for {
		n = num % 10
		num /= 10
		buf += n

		if num == 0 {

			if buf/10 == 0 {
				break
			}
			num = buf
			buf = 0
		}
	}
	return buf
}
