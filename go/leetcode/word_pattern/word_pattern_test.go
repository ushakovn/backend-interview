package word_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	pattern  string
	s        string
	expected bool
}

var testCases = []*testCase{
	{
		pattern:  "abba",
		s:        "dog cat cat dog",
		expected: true,
	},
	{
		pattern:  "abba",
		s:        "dog cat cat fish",
		expected: false,
	},
	{
		pattern:  "aaaa",
		s:        "dog cat cat dog",
		expected: false,
	},
}

func Test_wordPattern(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = wordPattern(testCase.pattern, testCase.s)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
