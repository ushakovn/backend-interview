package rotate_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	matrix   [][]int
	expected [][]int
}

var testCases = []*testCase{
	{
		matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		expected: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		},
	},
	{
		matrix: [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
		expected: [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		},
	},
}

func Test_rotate(t *testing.T) {
	for _, testCase := range testCases {
		rotate(testCase.matrix)

		if !assert.Equal(t, testCase.expected, testCase.matrix) {
			break
		}
	}
}
