package binary_insert

func searchInsert(nums []int, target int) int {
	var (
		mid   int
		count int
	)
	if count = len(nums); count == 0 || target <= nums[0] {
		return 0
	}
	if target > nums[count-1] {
		return count
	}
	var (
		left  int
		right = count - 1
	)
	for left <= right {
		mid = (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return left
}
