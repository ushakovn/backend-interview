package merge_sorted_array

import "sort"

func merge(nums1 []int, m int, nums2 []int, _ int) []int {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	return nums1
}
