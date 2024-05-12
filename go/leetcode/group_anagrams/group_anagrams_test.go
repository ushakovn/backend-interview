package group_anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	type testCase struct {
		strs     []string
		expected [][]string
	}
	testCases := []testCase{
		{
			strs: []string{
				"eat", "tea", "tan", "ate", "nat", "bat",
			},
			expected: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}
	for _, tc := range testCases {
		actual := groupAnagramsV1(tc.strs)

		assert.Equal(t, tc.expected, actual)
	}
}
