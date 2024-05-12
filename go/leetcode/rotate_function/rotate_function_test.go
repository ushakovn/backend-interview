package rotate_function

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
		nums:     []int{4, 3, 2, 6},
		expected: 26,
	},
	{
		nums:     []int{100},
		expected: 0,
	},
}

func Test_maxRotateFunction(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = maxRotateFunction(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
