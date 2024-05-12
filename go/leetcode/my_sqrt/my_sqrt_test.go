package my_sqrt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	x        int
	expected int
}

var testCases = []*testCase{
	{
		x:        4,
		expected: 2,
	},
	{
		x:        8,
		expected: 2,
	},
}

func Test_mySqrt(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = mySqrt(testCase.x)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
