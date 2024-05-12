package squares_sorted_array

func sortedSquares(nums []int) []int {
	count := len(nums)
	out := make([]int, count)

	leftPtr := 0
	rightPtr := count - 1

	for idx := 0; idx < count; idx++ {
		leftSqr := nums[leftPtr] * nums[leftPtr]
		rightSqr := nums[rightPtr] * nums[rightPtr]

		if leftSqr <= rightSqr {
			out[count-1-idx] = rightSqr
			rightPtr--
		} else {
			out[count-1-idx] = leftSqr
			leftPtr++
		}
	}
	return out
}
