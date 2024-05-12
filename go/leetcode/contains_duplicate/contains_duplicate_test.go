package contains_duplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected bool
}

var testCases = []*testCase{
	{
		nums:     []int{1, 2, 3, 1},
		expected: true,
	},
	{
		nums:     []int{1, 2, 3, 4},
		expected: false,
	},
	{
		nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		expected: true,
	},
}

func Test_containsDuplicate(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = containsDuplicate(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
