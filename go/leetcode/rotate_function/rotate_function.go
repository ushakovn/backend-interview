package rotate_function

func maxRotateFunction(nums []int) int {
	var (
		max int
		out int
	)
	out = f(nums)
	max = out

	for i := 1; i < len(nums); i++ {
		rotate(nums)

		if out = f(nums); out >= max {
			max = out
		}
	}
	return max
}

func f(nums []int) int {
	var out int

	for i, num := range nums {
		out += i * num
	}
	return out
}

func rotate(nums []int) {
	var (
		count int
		last  int
	)
	if count = len(nums); count == 0 {
		return
	}
	last = nums[count-1]

	for i := count - 1; i >= 0; i-- {
		if i == 0 {
			nums[i] = last
		} else {
			nums[i] = nums[i-1]
		}
	}
}
