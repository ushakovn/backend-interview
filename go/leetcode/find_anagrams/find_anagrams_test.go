package find_anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s        string
	p        string
	expected []int
}

var testCases = []*testCase{
	{
		s:        "cbaebabacd",
		p:        "abc",
		expected: []int{0, 6},
	},
	{
		s:        "abab",
		p:        "ab",
		expected: []int{0, 1, 2},
	},
}

func Test_findAnagrams(t *testing.T) {
	var actual []int

	for _, testCase := range testCases {
		actual = findAnagrams(testCase.s, testCase.p)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
