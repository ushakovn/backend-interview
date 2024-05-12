package rotate_array

func rotate(nums []int, k int) {
	for k != 0 {
		rot(nums)
		k--
	}
}

func rot(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var (
		ok bool
		n  int
	)
	for i := len(nums) - 1; i > 0; i-- {
		if !ok {
			n = nums[i]
			ok = true
		}
		nums[i] = nums[i-1]
	}
	nums[0] = n
}
