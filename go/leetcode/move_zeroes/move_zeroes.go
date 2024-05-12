package move_zeroes

func moveZeroes(nums []int) {
	out := make([]int, 0, len(nums))
	zeroes := make([]int, 0, len(nums))

	for _, num := range nums {
		if num == 0 {
			zeroes = append(zeroes, 0)
			continue
		}
		out = append(out, num)
	}
	out = append(out, zeroes...)

	copy(nums, out)
}
