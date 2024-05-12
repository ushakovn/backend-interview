package divide_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	dividend int
	divisor  int
	expected int
}

var testCases = []*testCase{
	{
		dividend: 10,
		divisor:  3,
		expected: 3,
	},
	{
		dividend: 7,
		divisor:  -3,
		expected: -2,
	},
	{
		dividend: 0,
		divisor:  0,
		expected: 0,
	},
	{
		dividend: -1,
		divisor:  1,
		expected: -1,
	},
	{
		dividend: -1,
		divisor:  -1,
		expected: 1,
	},
	{
		dividend: 8765432,
		divisor:  -1,
		expected: -8765432,
	},
	{
		dividend: -456677,
		divisor:  -1,
		expected: 456677,
	},
}

func Test_divide(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = divide(testCase.dividend, testCase.divisor)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
