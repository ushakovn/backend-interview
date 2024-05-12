package first_occurrence_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	haystack string
	needle   string
	expected int
}

var testCases = []*testCase{
	{
		haystack: "sadbutsad",
		needle:   "sad",
		expected: 0,
	},
	{
		haystack: "leetcode",
		needle:   "leeto",
		expected: -1,
	},
	{
		haystack: "mississippi",
		needle:   "issip",
		expected: 4,
	},
}

func Test_strStr(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = strStr(testCase.haystack, testCase.needle)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
