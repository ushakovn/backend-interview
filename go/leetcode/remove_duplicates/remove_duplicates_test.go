package remove_duplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	x        []int
	expected int
}

var testCases = []*testCase{
	{
		x:        []int{1, 1, 2},
		expected: 2,
	},
	{
		x:        []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		expected: 5,
	},
	{
		x:        []int{0},
		expected: 1,
	},
}

func Test_removeDuplicates(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = removeDuplicates(testCase.x)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
