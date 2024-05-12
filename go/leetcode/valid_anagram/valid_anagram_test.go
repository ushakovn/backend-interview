package valid_anagram

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
		s:        "anagram",
		t:        "nagaram",
		expected: true,
	},
	{
		s:        "rat",
		t:        "car",
		expected: false,
	},
	{
		s:        "a",
		t:        "b",
		expected: false,
	},
}

func Test_isAnagram(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = isAnagram(testCase.s, testCase.t)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
