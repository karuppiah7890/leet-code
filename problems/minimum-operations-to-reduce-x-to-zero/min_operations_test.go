package min_operations_test

import (
	"testing"

	min_operations "github.com/karuppiah7890/leet-code/problems/minimum-operations-to-reduce-x-to-zero"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums           []int
	x              int
	expectedOutput int
}

type TestCases []TestCase

func TestMinOperations(t *testing.T) {
	testCases := TestCases{
		{nums: []int{1, 1, 4, 2, 3}, x: 5, expectedOutput: 2},
		{nums: []int{5, 6, 7, 8, 9}, x: 4, expectedOutput: -1},
		{nums: []int{3, 2, 20, 1, 1, 3}, x: 10, expectedOutput: 5},
	}

	for _, testCase := range testCases {
		output := min_operations.MinOperations(testCase.nums, testCase.x)

		assert.Equal(t, testCase.expectedOutput, output)
	}
}
