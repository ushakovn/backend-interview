package product_except_self

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected []int
}

var testCases = []*testCase{
	{
		nums:     []int{1, 2, 3, 4},
		expected: []int{24, 12, 8, 6},
	},
	{
		nums:     []int{-1, 1, 0, -3, 3},
		expected: []int{0, 0, 9, 0, 0},
	},
}

func Test_productExceptSelf(t *testing.T) {
	var actual []int

	for _, testCase := range testCases {
		actual = productExceptSelf(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
