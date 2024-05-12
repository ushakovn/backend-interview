package product_except_self

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	mul := 1

	for idx := 0; idx < len(nums); idx++ {
		out[idx] = mul
		mul *= nums[idx]
	}
	mul = 1

	for idx := len(out) - 1; idx >= 0; idx-- {
		out[idx] *= mul
		mul *= nums[idx]
	}
	return out
}
