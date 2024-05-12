package intersection_two_arrays

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}

	for _, num1 := range nums1 {
		m[num1]++
	}
	var nums3 []int

	for _, num2 := range nums2 {
		if count, ok := m[num2]; ok && count > 0 {
			m[num2]--
			nums3 = append(nums3, num2)
		}
	}
	return nums3
}
