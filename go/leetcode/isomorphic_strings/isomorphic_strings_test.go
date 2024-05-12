package isomorphic_strings

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
		s:        "egg",
		t:        "add",
		expected: true,
	},
	{
		s:        "foo",
		t:        "bar",
		expected: false,
	},
	{
		s:        "paper",
		t:        "title",
		expected: true,
	},
	{
		s:        "badc",
		t:        "baba",
		expected: false,
	},
}

func Test_isIsomorphic(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = isIsomorphic(testCase.s, testCase.t)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
