package my_pow

func myPow(x float64, n int) float64 {
	if x > 100 || x < -100 || n == 0 || x == 0 || x == 1 {
		return 1
	}
	var (
		p  = 1.0
		nc = n
	)
	if nc < 0 {
		nc *= -1
	}
	for idx := 0; idx < nc; idx++ {
		if n > 0 {
			p *= x
		} else {
			p *= 1.0 / x
		}
		if p < -10000 || p > 10000 {
			return 1
		}
	}
	return p
}
