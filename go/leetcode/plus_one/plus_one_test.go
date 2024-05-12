package plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	digits   []int
	expected []int
}

var testCases = []*testCase{
	{
		digits:   []int{1, 2, 3},
		expected: []int{1, 2, 4},
	},
	{
		digits:   []int{4, 3, 2, 1},
		expected: []int{4, 3, 2, 2},
	},
	{
		digits:   []int{9},
		expected: []int{1, 0},
	},
	{
		digits:   []int{1, 9, 0, 9, 9, 9},
		expected: []int{1, 9, 1, 0, 0, 0},
	},
}

func Test_plusOne(t *testing.T) {
	var actual []int

	for _, testCase := range testCases {
		actual = plusOne(testCase.digits)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
