package remove_duplicates

func removeDuplicates(nums []int) int {
	count := len(nums)

	unique := make([]int, count)
	m := make(map[int]struct{}, count)
	var k int

	for _, num := range nums {
		if _, ok := m[num]; !ok {
			unique[k] = num
			m[num] = struct{}{}
			k++
		}
	}
	copy(nums, unique)

	return k
}
