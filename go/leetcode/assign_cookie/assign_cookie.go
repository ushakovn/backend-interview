package assign_cookie

func findContentChildren(g []int, s []int) int {
	var (
		assigned int
		used     = make([]bool, len(s))
	)
	sort(g)
	sort(s)

	for _, g := range g {
		for i, s := range s {
			if s >= g && !used[i] {
				assigned++
				used[i] = true
				break
			}
		}
	}
	return assigned
}

func sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
