package is_monotonic

func isMonotonic(in []int) bool {
	var (
		fall = true
		rise = true
	)
	for i := 0; i < len(in)-1; i++ {

		fall = fall && in[i] <= in[i+1]
		rise = rise && in[i] >= in[i+1]

		if !fall && !rise {
			return false
		}
	}
	return true
}
