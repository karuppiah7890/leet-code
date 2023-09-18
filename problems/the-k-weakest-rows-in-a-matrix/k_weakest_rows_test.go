package k_weakest_rows_test

import (
	"testing"

	k_weakest_rows "github.com/karuppiah7890/leet-code/problems/the-k-weakest-rows-in-a-matrix"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	matrix         [][]int
	k              int
	expectedOutput []int
}

type TestCases []TestCase

func TestKWeakestRows(t *testing.T) {
	testCases := TestCases{
		{
			matrix: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			k:              3,
			expectedOutput: []int{2, 0, 3},
		},
		{
			matrix: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			k:              2,
			expectedOutput: []int{0, 2},
		},
	}

	for _, testCase := range testCases {
		output := k_weakest_rows.KWeakestRows(testCase.matrix, testCase.k)
		assert.Equal(t, testCase.expectedOutput, output)
	}
}

func BenchmarkKWeakestRowsBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k_weakest_rows.KWeakestRowsBruteForce([][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		}, 3)
	}
}

func BenchmarkKWeakestRowsQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k_weakest_rows.KWeakestRowsQuick([][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		}, 3)
	}
}
