package length_of_last_word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s        string
	expected int
}

var testCases = []*testCase{
	{
		s:        "Hello World",
		expected: 5,
	},
	{
		s:        "   fly me   to   the moon  ",
		expected: 4,
	},
	{
		s:        "luffy is still joyboy",
		expected: 6,
	},
}

func Test_lengthOfLastWord(t *testing.T) {
	var actual int

	for _, testCase := range testCases {
		actual = lengthOfLastWord(testCase.s)

		if !assert.Equal(t, testCase.expected, actual) {
			break
		}
	}

}
