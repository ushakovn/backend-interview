package group_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s        []string
	expected [][]string
}

var testCases = []*testCase{
	{
		s:        []string{"ate", "eat", "tea", "nat", "tan", "bat"},
		expected: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
	},
}

func Test_groupWords(t *testing.T) {
	var actual [][]string

	for _, testCase := range testCases {
		actual = groupWordsV2(testCase.s)

		if assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
