package contains_duplicate_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	k        int
	expected bool
}

var testCases = []*testCase{
	{
		nums:     []int{1, 2, 3, 1},
		k:        3,
		expected: true,
	},
	{
		nums:     []int{1, 0, 1, 1},
		k:        1,
		expected: true,
	},
	{
		nums:     []int{1, 2, 3, 1, 2, 3},
		k:        2,
		expected: false,
	},
}

func Test_containsNearbyDuplicate(t *testing.T) {
	var actual bool

	for _, testCase := range testCases {
		actual = containsNearbyDuplicate(testCase.nums, testCase.k)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
