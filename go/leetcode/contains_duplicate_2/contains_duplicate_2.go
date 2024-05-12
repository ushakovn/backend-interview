package contains_duplicate_2

func containsNearbyDuplicate(nums []int, k int) bool {
	indexes := make(map[int]int, len(nums))

	for idx, num := range nums {
		if jdx, ok := indexes[num]; !ok || idx-jdx > k {
			indexes[num] = idx
			continue
		}
		return true
	}
	return false
}
