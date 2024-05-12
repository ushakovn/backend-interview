package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s        string
	expected bool
}

var testCases = []*testCase{
	{
		s:        "([([{}])])",
		expected: true,
	},
	{
		s:        "()[]{}",
		expected: true,
	},
	{
		s:        "()",
		expected: true,
	},
	{
		s:        "([([{])]])",
		expected: false,
	},
}

func Test_isValid(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = isValid(testCase.s)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}

}
