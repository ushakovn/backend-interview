package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	strs     []string
	expected string
}

var testCases = []*testCase{
	{
		strs:     []string{"flower", "flow", "flight"},
		expected: "fl",
	},
	{
		strs:     []string{"dog", "racecar", "car"},
		expected: "",
	},
}

func Test_longestCommonPrefix(t *testing.T) {
	var actual string

	for _, testCase := range testCases {
		actual = longestCommonPrefix(testCase.strs)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
