package max_sub_array

// read: https://en.wikipedia.org/wiki/Maximum_subarray_problem

func maxSubArray(nums []int) int {
	if count := len(nums); count == 0 {
		return 0
	}
	curSum := nums[0]
	maxSum := nums[0]

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for idx, num := range nums {
		if idx == 0 {
			continue
		}
		curSum = max(num, num+curSum)
		maxSum = max(curSum, maxSum)
	}
	return maxSum
}
