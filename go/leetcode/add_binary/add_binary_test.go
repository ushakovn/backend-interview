package add_binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	a        string
	b        string
	expected string
}

var testCases = []*testCase{
	{
		a:        "11",
		b:        "1",
		expected: "100",
	},
	{
		a:        "1010",
		b:        "1011",
		expected: "10101",
	},
}

func Test_addBinary(t *testing.T) {
	var actual string

	for _, testCase := range testCases {
		actual = addBinary(testCase.a, testCase.b)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}
}
