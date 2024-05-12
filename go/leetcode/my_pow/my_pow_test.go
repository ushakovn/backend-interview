package my_pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	x        float64
	n        int
	expected float64
}

var testCases = []*testCase{
	{
		x:        2.00000,
		n:        10,
		expected: 1024.00000,
	},
	{
		x:        2.00000,
		n:        -2,
		expected: 0.25000,
	},
	{
		x:        -2.00000,
		n:        2,
		expected: 4,
	},
}

func Test_myPow(t *testing.T) {
	var actual float64

	for _, testCase := range testCases {
		actual = myPow(testCase.x, testCase.n)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
