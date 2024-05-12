package single_number

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
		nums:     []int{2, 2, 1},
		expected: 1,
	},
	{
		nums:     []int{4, 1, 2, 1, 2},
		expected: 4,
	},
	{
		nums:     []int{1},
		expected: 1,
	},
	{
		nums:     []int{},
		expected: 0,
	},
}

func Test_singleNumber(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = singleNumber(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
