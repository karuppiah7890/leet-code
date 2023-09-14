package twosum_test

import (
	"testing"

	twosum "github.com/karuppiah7890/leet-code/problems/two-sum"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums           []int
	target         int
	expectedOutput []int
}

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{nums: []int{2, 7, 11, 15}, target: 9, expectedOutput: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expectedOutput: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expectedOutput: []int{0, 1}},
	}

	for _, testCase := range testCases {
		output := twosum.TwoSum(testCase.nums, testCase.target)
		assert.ElementsMatch(t, testCase.expectedOutput, output)
	}
}
