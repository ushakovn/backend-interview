package assign_cookie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	g        []int
	s        []int
	expected int
}

var testCases = []*testCase{
	{
		g:        []int{10, 9, 8, 7},
		s:        []int{10, 9, 8, 7},
		expected: 4,
	},
	{
		g:        []int{1, 2, 3},
		s:        []int{1, 1},
		expected: 1,
	},
}

func Test_findContentChildren(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = findContentChildren(testCase.g, testCase.s)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
