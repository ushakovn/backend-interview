package divide_integers

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}
	var (
		divideCount int
		dividendNeg bool
		divisorNeg  bool
	)
	if dividend < 0 {
		dividend = swapSign(dividend)
		dividendNeg = true
	}
	if divisor < 0 {
		divisor = swapSign(divisor)
		divisorNeg = true
	}
	for dividend >= divisor {
		dividend -= divisor
		divideCount++
	}
	if dividendNeg != divisorNeg {
		divideCount = swapSign(divideCount)
	}
	return divideCount
}

func swapSign(num int) int {
	buf := num
	num -= buf
	num -= buf
	return num
}
