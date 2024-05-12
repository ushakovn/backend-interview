package max_sub_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected int
}

var testCases = []*testCase{
	{
		nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	},
	{
		nums:     []int{5, 4, -1, 7, 8},
		expected: 23,
	},
	{
		nums:     []int{1},
		expected: 1,
	},
}

func Test_maxSubArray(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = maxSubArray(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
