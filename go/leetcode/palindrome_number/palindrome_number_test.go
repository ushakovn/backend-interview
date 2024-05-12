package palindrome_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	x        int
	expected bool
}

var testCases = []*testCase{
	{
		x:        121,
		expected: true,
	},
	{
		x:        -121,
		expected: false,
	},
	{
		x:        10,
		expected: false,
	},
}

func Test_isPalindrome(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = isPalindrome(testCase.x)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
