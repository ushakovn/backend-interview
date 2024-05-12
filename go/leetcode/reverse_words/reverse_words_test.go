package reverse_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s        string
	expected string
}

var testCases = []*testCase{
	{
		s:        "Let's take LeetCode contest",
		expected: "s'teL ekat edoCteeL tsetnoc",
	},
	{
		s:        "God Ding",
		expected: "doG gniD",
	},
	{
		s:        "   ",
		expected: "   ",
	},
}

func Test_reverseWords(t *testing.T) {
	var actual string

	for _, testCase := range testCases {
		actual = reverseWords(testCase.s)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
