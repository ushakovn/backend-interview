package majority_element

func majorityElement(nums []int) int {
	m := map[int]int{}
	total := len(nums)

	for _, num := range nums {
		m[num]++
	}
	for _, num := range nums {
		if count := m[num]; count > total/2 {
			return num
		}
	}
	return 0
}
