package majority_element

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
		nums:     []int{3, 2, 3},
		expected: 3,
	},
	{
		nums:     []int{2, 2, 1, 1, 1, 2, 2},
		expected: 2,
	},
	{
		nums:     []int{},
		expected: 0,
	},
}

func Test_majorityElement(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = majorityElement(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
