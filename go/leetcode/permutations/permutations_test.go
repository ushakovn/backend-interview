package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected [][]int
}

var testCases = []*testCase{
	{
		nums: []int{1, 2, 3},
		expected: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		},
	},
	{
		nums: []int{0, 1},
		expected: [][]int{
			{0, 1},
			{1, 0},
		},
	},
	{
		nums: []int{1},
		expected: [][]int{
			{1},
		},
	},
}

func Test_permute(t *testing.T) {
	var actual [][]int

	for _, testCase := range testCases {
		actual = permute(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
