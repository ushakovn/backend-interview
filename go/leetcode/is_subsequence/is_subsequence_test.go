package is_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s        string
	t        string
	expected bool
}

var testCases = []*testCase{
	{
		s:        "abc",
		t:        "ahbgdc",
		expected: true,
	},
	{
		s:        "axc",
		t:        "ahbgdc",
		expected: false,
	},
}

func Test_isSubsequence(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = isSubsequence(testCase.s, testCase.t)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
