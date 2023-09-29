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
		{nums: []int{}, x: 4, expectedOutput: -1},
		{nums: []int{4}, x: 4, expectedOutput: 1},
		{nums: []int{5}, x: 4, expectedOutput: -1},
		{nums: []int{3}, x: 4, expectedOutput: -1},
		{nums: []int{5, 5}, x: 4, expectedOutput: -1},
		{nums: []int{4, 4}, x: 4, expectedOutput: 1},
		{nums: []int{5, 4}, x: 4, expectedOutput: 1},
		{nums: []int{4, 5}, x: 4, expectedOutput: 1},
		{nums: []int{4, 3, 5}, x: 4, expectedOutput: 1},
		{nums: []int{5, 3, 4}, x: 4, expectedOutput: 1},
		{nums: []int{5, 6, 7, 8, 9}, x: 4, expectedOutput: -1},
		{nums: []int{2, 2}, x: 4, expectedOutput: 2},
		{nums: []int{2, 2, 2}, x: 6, expectedOutput: 3},
		{nums: []int{10, 2, 2, 2}, x: 6, expectedOutput: 3},
		{nums: []int{2, 2, 2, 10}, x: 6, expectedOutput: 3},
		// Consider a case where there are multiple ways to make x as 0, then
		// choose the minimum number of operations to make x as 0
		{nums: []int{1, 1, 4, 2, 3}, x: 5, expectedOutput: 2}, // 1,1,3 or 3,2
		{nums: []int{3, 2, 4, 1, 1}, x: 5, expectedOutput: 2}, // 1,1,3 or 3,2
		{nums: []int{3, 2, 20, 1, 1, 3}, x: 10, expectedOutput: 5},
		{nums: []int{5, 2, 5}, x: 6, expectedOutput: -1},
		{nums: []int{1241, 8769, 9151, 3211, 2314,
			8007, 3713, 5835, 2176, 8227, 5251,
			9229, 904, 1899, 5513, 7878, 8663,
			3804, 2685, 3501, 1204, 9742, 2578,
			8849, 1120, 4687, 5902, 9929, 6769,
			8171, 5150, 1343, 9619, 3973, 3273,
			6427, 47, 8701, 2741, 7402, 1412,
			2223, 8152, 805, 6726, 9128, 2794,
			7137, 6725, 4279, 7200, 5582, 9583,
			7443, 6573, 7221, 1423, 4859, 2608,
			3772, 7437, 2581, 975, 3893, 9172,
			3, 3113, 2978, 9300, 6029, 4958, 229,
			4630, 653, 1421, 5512, 5392, 7287, 8643,
			4495, 2640, 8047, 7268, 3878, 6010, 8070,
			7560, 8931, 76, 6502, 5952, 4871, 5986,
			4935, 3015, 8263, 7497, 8153, 384, 1136}, x: 894887480, expectedOutput: -1},
	}

	for _, testCase := range testCases {
		output := min_operations.MinOperations(testCase.nums, testCase.x)

		assert.Equal(t, testCase.expectedOutput, output, "Input: nums: %v, x: %v", testCase.nums, testCase.x)
	}
}
