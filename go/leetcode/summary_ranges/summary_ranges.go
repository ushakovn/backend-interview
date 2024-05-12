package summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	count := len(nums)

	if count == 0 {
		return nil
	}
	ranges := make([]string, 0, count)

	var (
		start int
		prev  int
		rang  string
	)
	push := func() {
		if prev-start > 0 {
			rang = fmt.Sprint(start, "->", prev)
		} else {
			rang = fmt.Sprint(start)
		}
		ranges = append(ranges, rang)
	}
	for idx, num := range nums {
		if idx == 0 {

			start = num
			prev = num

			continue
		}
		if num-prev > 1 {
			push()

			start = num
			prev = num
		}
		prev = num
	}
	push()

	return ranges
}
