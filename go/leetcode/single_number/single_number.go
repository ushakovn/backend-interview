package single_number

func singleNumber(nums []int) int {
	counts := make(map[int]int, len(nums)/2+1)

	for _, num := range nums {
		counts[num]++
	}
	var sn int

	for _, num := range nums {
		if count := counts[num]; count == 1 {
			sn = num
			break
		}
	}
	return sn
}
