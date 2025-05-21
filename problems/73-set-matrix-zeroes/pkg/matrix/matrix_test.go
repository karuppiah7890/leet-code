package matrix

import (
	"fmt"
	"testing"
)

type TestCaseInput struct {
	matrix [][]int
}

type TestCaseOutput struct {
	matrix [][]int
}

type TestCase struct {
	input          TestCaseInput
	expectedOutput TestCaseOutput
}

type TestCases []TestCase

func TestSetZeroes(t *testing.T) {
	testCases := getTestCases()

	for index, testCase := range testCases {
		testName := fmt.Sprintf("TestSetZeroes-%d", index)
		t.Run(testName, func(t *testing.T) {
			inputBeforeProcessing, err := cloneMatrix(testCase.input.matrix)
			if err != nil {
				t.Error(err)
				return
			}

			setZeroes(testCase.input.matrix)

			output := testCase.input.matrix

			if !isMatrixEqual(output, testCase.expectedOutput.matrix) {
				t.Errorf("error in test case index: %d. Input (Before Processing): %v. Output: %v. Expected Output: %v", index, inputBeforeProcessing, output, testCase.expectedOutput.matrix)
			}
		})
	}

}

func cloneMatrix(matrix [][]int) ([][]int, error) {
	clonedMatrix := make([][]int, len(matrix))
	n := copy(clonedMatrix, matrix)
	if n != len(matrix) {
		return nil, fmt.Errorf("error occurred while doing copy operation on matrix input. Input: %v. Copied Input: %v", matrix, clonedMatrix)
	}
	return clonedMatrix, nil
}

func isMatrixEqual(matrix [][]int, anotherMatrix [][]int) bool {
	if len(matrix) != len(anotherMatrix) {
		return false
	}

	for rowIndex, row := range matrix {
		if !isArrayEqual(row, anotherMatrix[rowIndex]) {
			return false
		}
	}

	return true
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
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			expectedOutput: TestCaseOutput{
				matrix: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 0, 1},
				},
			},
		},
		TestCase{
			input: TestCaseInput{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			expectedOutput: TestCaseOutput{
				matrix: [][]int{
					{0, 0, 0, 0},
					{0, 4, 5, 0},
					{0, 3, 1, 0},
				},
			},
		},
	}
}
