package search_insert_position

func searchInsert(nums []int, target int) int {
	var (
		count int
		index int
	)
	if count = len(nums); count <= 0 || target <= nums[0] {
		return 0
	}
	if target > nums[count-1] {
		return count
	}
	for idx := 0; idx < count-1; idx++ {
		if target >= nums[idx] && target <= nums[idx+1] {
			if target == nums[idx] {
				index = idx
			}
			index = idx + 1
			break
		}
	}
	return index
}
