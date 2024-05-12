package reverse_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	x        int
	expected int
}

var testCases = []*testCase{
	{
		x:        123,
		expected: 321,
	},
	{
		x:        -123,
		expected: -321,
	},
	{
		x:        120,
		expected: 21,
	},
}

func Test_reverse(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = reverse(testCase.x)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
