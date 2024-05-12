package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

var testCases = []*testCase{
	{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
}

var testFunctions = []func([]int, int) []int{
	twoSumV2,
	twoSum,
}

func Test_twoSum_twoSumWithMap(t *testing.T) {
	actual := make([]int, 2)

	for _, testFunc := range testFunctions {
		for _, testCase := range testCases {
			actual = testFunc(testCase.nums, testCase.target)

			if !assert.Equal(t, testCase.expected, actual) {
				return
			}
		}
	}
}
