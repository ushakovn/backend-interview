package max_product_sub_array

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
		nums:     []int{2, 3, -2, 4},
		expected: 6,
	},
	{
		nums:     []int{-2, 0, -1},
		expected: 0,
	},
}

func Test_maxSubArray(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = maxProduct(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
