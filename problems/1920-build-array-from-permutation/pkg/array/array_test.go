package array

import (
	"fmt"
	"testing"
)

type TestCaseInput struct {
	nums []int
}

type TestCaseOutput struct {
	ans []int
}

type TestCase struct {
	input          TestCaseInput
	expectedOutput TestCaseOutput
}

type TestCases []TestCase

func TestBuildArray(t *testing.T) {
	testCases := getTestCases()

	for index, testCase := range testCases {
		testName := fmt.Sprintf("TestBuildArray-%d", index)
		t.Run(testName, func(t *testing.T) {
			inputBeforeProcessing, err := cloneArray(testCase.input.nums)
			if err != nil {
				t.Error(err)
				return
			}

			output := buildArray(testCase.input.nums)

			if !isArrayEqual(output, testCase.expectedOutput.ans) {
				t.Errorf("error in test case index: %d. Input (Before Processing): %v. Output: %v. Expected Output: %v", index, inputBeforeProcessing, output, testCase.expectedOutput.ans)
			}
		})
	}

}

func cloneArray(array []int) ([]int, error) {
	clonedArray := make([]int, len(array))
	n := copy(clonedArray, array)
	if n != len(array) {
		return nil, fmt.Errorf("error occurred while doing copy operation on array input. Input: %v. Copied Input: %v", array, clonedArray)
	}
	return clonedArray, nil
}

func isArrayEqual(array []int, anotherArray []int) bool {
	if len(array) != len(anotherArray) {
		return false
	}

	for index, element := range array {
		if element != anotherArray[index] {
			return false
		}
	}

	return true
}

func getTestCases() TestCases {
	return TestCases{
		TestCase{
			input: TestCaseInput{
				nums: []int{0},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{0},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{0, 1, 2},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{0, 1, 2},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{0, 2, 1, 5, 3, 4},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{0, 1, 2, 4, 5, 3},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{0, 11, 9, 4, 7, 10, 5, 8, 6, 3, 2, 1},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{0, 1, 3, 7, 8, 2, 10, 6, 5, 4, 9, 11},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{0, 11, 9, 4, 7, 10, 5, 8, 6, 3, 1, 2},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{0, 2, 3, 7, 8, 1, 10, 6, 5, 4, 11, 9},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{0, 11, 9, 4, 7, 10, 5, 8, 6, 3, 2, 17, 12, 13, 14,
					15, 1, 16},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{0, 17, 3, 7, 8, 2, 10, 6, 5, 4, 9, 16, 12, 13, 14, 15, 11, 1},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{10, 0, 9, 1, 7, 11, 5, 8, 6, 3, 4, 2},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{4, 10, 3, 0, 8, 2, 11, 6, 5, 1, 7, 9},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{6, 11, 9, 4, 7, 10, 5, 8, 0, 3, 2, 17, 12, 13, 14, 15, 1, 16},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{5, 17, 3, 7, 8, 2, 10, 0, 6, 4, 9, 16, 12, 13, 14, 15, 11, 1},
			},
		},
		TestCase{
			input: TestCaseInput{
				nums: []int{6, 11, 9, 4, 7, 10, 0, 8, 5, 3, 2, 17, 12, 13, 14, 15, 1, 16},
			},
			expectedOutput: TestCaseOutput{
				ans: []int{0, 17, 3, 7, 8, 2, 6, 5, 10, 4, 9, 16, 12, 13, 14, 15, 11, 1},
			},
		},
	}
}
