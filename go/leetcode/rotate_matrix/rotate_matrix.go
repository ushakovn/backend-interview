package rotate_matrix

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	out := make([][]int, len(matrix))
	k := 0

	for i := 0; k < len(matrix) && i < len(matrix[k]); i++ {
		buf := make([]int, len(matrix))

		for j := len(matrix) - 1; j >= 0; j-- {
			buf[len(matrix)-j-1] = matrix[j][i]
		}
		out[i] = buf
		k++
	}
	copy(matrix, out)
}
