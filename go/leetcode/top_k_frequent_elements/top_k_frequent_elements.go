package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}
	buf := make([][]int, len(nums))

	for num, count := range m {
		buf[count-1] = append(buf[count-1], num)
	}
	top := make([]int, k)
	d := 0

	for i := len(buf) - 1; i >= 0 && k > 0; i-- {
		for j := 0; j < len(buf[i]) && k > 0; j++ {
			top[d] = buf[i][j]
			d++
			k--
		}
	}
	return top
}
