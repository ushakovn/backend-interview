package permutations

func permute(nums []int) [][]int {
	var (
		perms [][]int
		perm  []int
	)
	check := make(map[int]struct{}, len(nums))

	backtrack(&perms, nums, perm, check)

	return perms
}

func backtrack(perms *[][]int, nums []int, perm []int, check map[int]struct{}) {
	if len(perm) == len(nums) {
		cp := make([]int, len(perm))

		copy(cp, perm)
		*perms = append(*perms, cp)

		return
	}
	for idx := 0; idx < len(nums); idx++ {
		num := nums[idx]

		if _, ok := check[num]; !ok {
			check[num] = struct{}{}
			perm = append(perm, num)

			backtrack(perms, nums, perm, check)

			delete(check, num)
			perm = perm[:len(perm)-1]
		}
	}
}
