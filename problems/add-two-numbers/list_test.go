package list_test

import (
	"fmt"
	"testing"

	list "github.com/karuppiah7890/leet-code/problems/add-two-numbers"
)

type NewTestCase struct {
	digits []int
}

type NewTestCases []NewTestCase

func TestNew(t *testing.T) {
	testCases := NewTestCases{
		{
			digits: nil,
		},
		{
			digits: []int{},
		},
		{
			digits: []int{1},
		},
		{
			digits: []int{1, 2},
		},
		{
			digits: []int{1, 2, 3, 4},
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", index+1), func(subT *testing.T) {
			digits := testCase.digits
			l := list.New(digits)
			var index int
			node := l

			for index = 0; index < len(digits) && node != nil; index++ {
				if digits[index] != node.Val {
					t.Errorf("expected the corresponding digits at index %d to be equal but they were not. digit from array: %v. digit from list: %v", index, digits[index], node.Val)
				}

				node = node.Next
			}

			if node != nil {
				subT.Errorf("expected to have traversed through whole list but didn't traverse through whole list")
			}

			if index != len(digits) {
				subT.Errorf("expected to have traversed through all digits but didn't traverse through all digits")
			}
		})
	}
}

type AddTestCase struct {
	list                 []int
	anotherList          []int
	expectedOutputDigits []int
}

type AddTestCases []AddTestCase

func TestAdd(t *testing.T) {
	testCases := AddTestCases{
		{
			list:                 []int{0},
			anotherList:          []int{0},
			expectedOutputDigits: []int{0},
		},
		{
			list:                 []int{1},
			anotherList:          []int{2},
			expectedOutputDigits: []int{3},
		},
		{
			list:                 []int{7},
			anotherList:          []int{2},
			expectedOutputDigits: []int{9},
		},
		{
			list:                 []int{7},
			anotherList:          []int{3},
			expectedOutputDigits: []int{0, 1},
		},
		// {
		// 	list:           []int{2, 4, 3},
		// 	anotherList:    []int{5, 6, 4},
		// 	expectedOutputDigits: []int{7, 0, 8},
		// },
		// {
		// 	list:           []int{9, 9, 9, 9, 9, 9, 9},
		// 	anotherList:    []int{9, 9, 9, 9},
		// 	expectedOutputDigits: []int{8, 9, 9, 9, 0, 0, 0, 1},
		// },
	}

	for _, testCase := range testCases {
		l1 := list.New(testCase.list)
		l2 := list.New(testCase.anotherList)

		output := l1.Add(l2)

		expectedOutput := list.New(testCase.expectedOutputDigits)

		if !output.Equals(expectedOutput) {
			t.Errorf("expected output to be something but it was not. actual: %v, expected: %v", output.Digits(), testCase.expectedOutputDigits)
		}
	}
}

func TestDigits(t *testing.T) {
	digits := []int{1, 2, 3, 4}
	l := list.New(digits)

	assertIntArrayEqual(t, l.Digits(), digits)
}

func assertIntArrayEqual(t *testing.T, actual []int, expected []int) {
	if len(actual) != len(expected) {
		t.Errorf("expected the two integer arrays to be equal but they were not. Their sizes are different. Actual: %v. Expected: %v", len(actual), len(expected))
	}

	for index := 0; index < len(actual); index++ {
		if actual[index] != expected[index] {
			t.Errorf("expected the two integers at index %d to be equal but they were not. Actual: %v. Expected: %v", index, actual[index], expected[index])
			return
		}
	}
}

func TestEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		l1 := list.New([]int{1, 2, 3, 4})
		l2 := list.New([]int{1, 2, 3, 4})

		if !l1.Equals(l2) {
			t.Errorf("expected the two lists to be equal but they were not")
		}

		if !l2.Equals(l1) {
			t.Errorf("expected the two lists to be equal but they were not")
		}
	})

	t.Run("unequal", func(t *testing.T) {
		t.Run("different sized lists", func(t *testing.T) {
			l1 := list.New([]int{1, 2})
			l2 := list.New([]int{1, 2, 3, 4})

			if l1.Equals(l2) {
				t.Errorf("expected the two lists to be unequal but they were equal")
			}

			if l2.Equals(l1) {
				t.Errorf("expected the two lists to be unequal but they were equal")
			}
		})

		t.Run("different values in lists", func(t *testing.T) {
			l1 := list.New([]int{1, 2})
			l2 := list.New([]int{1, 3})

			if l1.Equals(l2) {
				t.Errorf("expected the two lists to be unequal but they were equal")
			}

			if l2.Equals(l1) {
				t.Errorf("expected the two lists to be unequal but they were equal")
			}
		})
	})
}
