package max_profit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	prices   []int
	expected int
}

var testCases = []*testCase{
	{
		prices:   []int{7, 1, 5, 3, 6, 4},
		expected: 5,
	},
	{
		prices:   []int{7, 6, 4, 3, 1},
		expected: 0,
	},
}

var testFunctions = []func([]int) int{
	maxProfitV2,
	maxProfit,
}

func Test_maxProfit(t *testing.T) {
	var actual int

	for _, testFunc := range testFunctions {
		for _, testCase := range testCases {
			actual = testFunc(testCase.prices)

			if !assert.Equal(t, testCase.expected, actual) {
				return
			}
		}
	}
}
