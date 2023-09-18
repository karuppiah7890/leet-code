package twosum_test

import (
	"testing"

	twosum "github.com/karuppiah7890/leet-code/problems/two-sum"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums           []int32
	target         int32
	expectedOutput []int32
}

func TestTwoSum(t *testing.T) {
	testCases := []TestCase{
		{nums: []int32{2, 7, 11, 15}, target: 9, expectedOutput: []int32{0, 1}},
		{nums: []int32{3, 2, 4}, target: 6, expectedOutput: []int32{1, 2}},
		{nums: []int32{3, 3}, target: 6, expectedOutput: []int32{0, 1}},
		{nums: []int32{0, 4, 3, 0}, target: 0, expectedOutput: []int32{0, 3}},
		{nums: []int32{-3,4,3,90}, target: 0, expectedOutput: []int32{0, 2}},
	}

	for _, testCase := range testCases {
		output := twosum.TwoSum(testCase.nums, testCase.target)
		assert.ElementsMatch(t, testCase.expectedOutput, output)
	}
}

func BenchmarkTwoSumBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twosum.TwoSumBruteForce([]int32{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twosum.TwoSumQuick([]int32{2, 7, 11, 15}, 9)
	}
}
