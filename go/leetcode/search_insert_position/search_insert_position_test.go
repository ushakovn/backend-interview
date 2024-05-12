package search_insert_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	target   int
	expected int
}

var testCases = []*testCase{
	{
		nums:     []int{1, 3, 5, 6},
		target:   5,
		expected: 2,
	},
	{
		nums:     []int{1, 3, 5, 6},
		target:   2,
		expected: 1,
	},
	{
		nums:     []int{1, 3, 5, 6},
		target:   7,
		expected: 4,
	},
}

func Test_searchInsert(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = searchInsert(testCase.nums, testCase.target)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
