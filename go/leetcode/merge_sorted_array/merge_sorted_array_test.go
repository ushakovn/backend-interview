package merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
}

var testCases = []*testCase{
	{
		nums1:    []int{1, 2, 3, 0, 0, 0},
		m:        3,
		nums2:    []int{2, 5, 6},
		n:        3,
		expected: []int{1, 2, 2, 3, 5, 6},
	},
	{
		nums1:    []int{1},
		m:        1,
		nums2:    []int{},
		n:        0,
		expected: []int{1},
	},
}

func Test_merge(t *testing.T) {
	var actual []int

	for _, testCase := range testCases {
		actual = merge(testCase.nums1, testCase.m, testCase.nums2, testCase.n)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
