package max_product_sub_array

// read: https://en.wikipedia.org/wiki/Maximum_subarray_problem

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxProd := nums[0]
	curProd := nums[0]

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
		curProd = max(num, num*curProd)
		maxProd = max(curProd, maxProd)
	}
	return maxProd
}
