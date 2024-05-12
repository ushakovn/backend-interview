package my_sqrt

func mySqrt(x int) int {
	xf := float64(x)
	x0 := float64(pickX0(x))

	x1 := newtonSqrt(x0, xf, 1e-1)

	return int(x1)
}

func newtonSqrt(x0, x, eps float64) float64 {
	var (
		x1  float64
		dif float64
	)
	for {
		x1 = 0.5 * (x0 + x/x0)

		if dif = x1 - x0; dif >= 0 && dif <= eps {
			break
		}
		if dif = x0 - x1; dif >= 0 && dif <= eps {
			break
		}
		x0 = x1
	}
	return x1
}

func pickX0(x int) int {
	var x0 int

	for {
		if x = x / 10; x == 0 {
			break
		}
		x0 += 10
	}
	if x0 == 0 {
		x0 = 1
	}
	return x0
}
