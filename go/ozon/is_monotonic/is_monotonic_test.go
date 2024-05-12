package is_monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	in       []int
	expected bool
}

var testCases = []*testCase{
	{
		in:       []int{1, 7},
		expected: true,
	},
	{
		in:       []int{1, 1},
		expected: true,
	},
	{
		in:       []int{23, 5, 23},
		expected: false,
	},
}

func Test_isMonotonic(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = isMonotonic(testCase.in)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
