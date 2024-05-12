package two_sum

func twoSumV2(nums []int, target int) []int {
	differences := make(map[int]int, len(nums))
	indices := make([]int, 0, 2)

	for jdx, num := range nums {
		diff := target - num

		if idx, ok := differences[num]; ok {
			indices = append(indices, idx, jdx)
			break
		}
		differences[diff] = jdx
	}
	return indices
}

func twoSum(nums []int, target int) []int {
	indices := make([]int, 0, 2)
	var found bool

	for idx := 0; idx < len(nums)-1 && !found; idx++ {
		for jdx := idx + 1; jdx < len(nums) && !found; jdx++ {
			if nums[idx]+nums[jdx] == target {
				indices = append(indices, idx, jdx)

				found = true
				break
			}
		}
	}
	return indices
}
