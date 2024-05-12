package add_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	num      int
	expected int
}

var testCases = []*testCase{
	{
		num:      38,
		expected: 2,
	},
	{
		num:      0,
		expected: 0,
	},
}

func Test_addDigits(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = addDigits(testCase.num)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}

}
