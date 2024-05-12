package summary_ranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected []string
}

var testCases = []*testCase{
	{
		nums:     []int{0, 1, 2, 4, 5, 7},
		expected: []string{"0->2", "4->5", "7"},
	},
	{
		nums:     []int{0, 2, 3, 4, 6, 8, 9},
		expected: []string{"0", "2->4", "6", "8->9"},
	},
}

func Test_summaryRanges(t *testing.T) {
	var actual []string

	for _, testCase := range testCases {
		actual = summaryRanges(testCase.nums)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
