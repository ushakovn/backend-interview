package rotate_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	k        int
	expected []int
}

var testCases = []*testCase{
	{
		nums:     []int{1, 2, 3, 4, 5, 6, 7},
		k:        3,
		expected: []int{5, 6, 7, 1, 2, 3, 4},
	},
	{
		nums:     []int{-1, -100, 3, 99},
		k:        2,
		expected: []int{3, 99, -1, -100},
	},
}

func Test_rotate(t *testing.T) {
	for _, testCase := range testCases {
		rotate(testCase.nums, testCase.k)

		if !assert.Equal(t, testCase.expected, testCase.nums) {
			break
		}
	}
}
