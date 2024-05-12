package reverse_words_2

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
		s:        "the sky is blue",
		expected: "blue is sky the",
	},
	{
		s:        "  hello world  ",
		expected: "world hello",
	},
	{
		s:        "a good   example",
		expected: "example good a",
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
