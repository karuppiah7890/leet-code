package domino

import "testing"

type TestCaseInput struct {
	tops    []int
	bottoms []int
}

type TestCaseOutput int

type TestCase struct {
	input          TestCaseInput
	expectedOutput TestCaseOutput
}

type TestCases []TestCase

func TestMinDominoRotations(t *testing.T) {
	testCases := TestCases{
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 3},
				bottoms: []int{1, 4},
			},
			expectedOutput: -1,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 2, 3},
				bottoms: []int{1, 1, 4},
			},
			expectedOutput: -1,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 2},
				bottoms: []int{1, 1},
			},
			expectedOutput: 0,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 2},
				bottoms: []int{1, 5},
			},
			expectedOutput: 0,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 5},
				bottoms: []int{5, 5},
			},
			expectedOutput: 0,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 5},
				bottoms: []int{1, 2},
			},
			expectedOutput: 1,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 5, 1},
				bottoms: []int{1, 2, 2},
			},
			expectedOutput: 1,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 5, 1, 2},
				bottoms: []int{1, 2, 2, 5},
			},
			expectedOutput: 1,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{5, 2, 1, 2},
				bottoms: []int{1, 2, 2, 5},
			},
			expectedOutput: -1,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 1, 2, 4, 2, 2},
				bottoms: []int{5, 2, 6, 2, 3, 2},
			},
			expectedOutput: 2,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{2, 1, 2, 4, 5, 2},
				bottoms: []int{5, 2, 6, 2, 3, 2},
			},
			expectedOutput: 2,
		},
		TestCase{
			input: TestCaseInput{
				tops:    []int{3, 5, 1, 2, 3},
				bottoms: []int{3, 6, 3, 3, 4},
			},
			expectedOutput: -1,
		},
	}
}
