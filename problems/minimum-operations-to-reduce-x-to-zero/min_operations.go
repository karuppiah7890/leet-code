package min_operations

import "fmt"

// It doesn't matter if one memoizes start index, end index AND the number x' (after operations)
// or if one memoizes start index and end index alone.
// Test this out.
// So, if both are the same, then it's better to not store x'.
// And how are both same? Well, theoretically, if start index
// and end index are fixed, then input x' will be fixed too
// because x' would be formed by subtracting all the numbers
// on the left of the start index from x and also subtracting
//  all the numbers on the right of the end index from x.
// so, x' is basically a function of start index and end index.
// given start index and end index, one can easily compute
// and tell x'. so it doesn't make sense to store it or use it
// in any way. it's redundant and unnecessary to store x'

type minOperationsFinder struct {
	resultsStore map[string]int
}

func MinOperations(nums []int, x int) int {
	// array of size 0
	if len(nums) == 0 {
		return -1
	}

	// array of size 1
	if len(nums) == 1 {
		if nums[0] == x {
			// 1 operation to subtract the single number from x
			return 1
		} else {
			// here (nums[0] > x) or (nums[0] < x).
			// hence no operations can be done to
			// make x as 0 by subtracting num[0] from it
			return -1
		}
	}

	m := minOperationsFinder{
		resultsStore: make(map[string]int),
	}

	result := m.findMinOperations(nums, x, 0, len(nums)-1)

	// fmt.Printf("%v\n", m.resultsStore)

	return result
}

func key(startIndex int, endIndex int) string {
	return fmt.Sprintf("%d,%d", startIndex, endIndex)
}

func (m *minOperationsFinder) getResultFromStore(startIndex int, endIndex int) (int, bool) {
	key := key(startIndex, endIndex)
	result, ok := m.resultsStore[key]
	if !ok {
		return -1, false
	}

	return result, true
}

func (m *minOperationsFinder) setResultInStore(startIndex int, endIndex int, result int) {
	key := key(startIndex, endIndex)
	existingResult, ok := m.resultsStore[key]
	if ok {
		panic(fmt.Sprintf("error occurred while trying to store result in store: key '%s' has an existing result %d. trying to overwrite with result %d", key, existingResult, result))
	}

	m.resultsStore[key] = result
}

// the start index and end index are both inclusive indices
func (m *minOperationsFinder) findMinOperations(nums []int, x int, startIndex int, endIndex int) int {
	// sub array of size 1
	if startIndex == endIndex {
		if nums[startIndex] == x {
			// 1 operation to subtract the single number from x
			return 1
		} else {
			// here (nums[startIndex] > x) or (nums[startIndex] < x).
			// hence no operations can be done to
			// make x as 0 by subtracting num[startIndex] from it
			return -1
		}
	}

	first := nums[startIndex]
	last := nums[endIndex]

	if first == x || last == x {
		// the first or last number could be the same as x.
		// or both the first and last number could be the same as x.
		// In both cases, just 1 operation to subtract the single number from x.
		return 1
	}
	if first > x && last > x {
		return -1
	}

	numberOfOperationsUsingFirstNumber := -1
	numberOfOperationsUsingLastNumber := -1
	// the first and/ the last number is less than x

	// remove first number and follow that path and try to make x as 0
	if result, ok := m.getResultFromStore(startIndex+1, endIndex); ok {
		numberOfOperationsUsingFirstNumber = result
	} else {
		numberOfOperationsUsingFirstNumber = m.findMinOperations(nums, x-first, startIndex+1, endIndex)
		m.setResultInStore(startIndex+1, endIndex, numberOfOperationsUsingFirstNumber)
	}

	// remove last number and follow that path and try to make x as 0

	if result, ok := m.getResultFromStore(startIndex, endIndex-1); ok {
		numberOfOperationsUsingLastNumber = result
	} else {
		numberOfOperationsUsingLastNumber = m.findMinOperations(nums, x-last, startIndex, endIndex-1)
		m.setResultInStore(startIndex, endIndex-1, numberOfOperationsUsingLastNumber)
	}

	// if both first and last number removal path does not
	// lead x to become 0 on subtracting the numbers, then
	// return -1
	if numberOfOperationsUsingFirstNumber == -1 && numberOfOperationsUsingLastNumber == -1 {
		return -1
	}

	// if any one - first or last number, along with
	// more numbers can be subtracted from x to make it 0
	// then return the number of operations.
	if numberOfOperationsUsingFirstNumber == -1 {
		return numberOfOperationsUsingLastNumber + 1
	}
	if numberOfOperationsUsingLastNumber == -1 {
		return numberOfOperationsUsingFirstNumber + 1
	}

	// if both first number and last number path can be used to
	// make x as 0, then return the minimum of the two
	// Note: Below logic works even if both the values are equal, which is
	// also possible
	if numberOfOperationsUsingFirstNumber < numberOfOperationsUsingLastNumber {
		return numberOfOperationsUsingFirstNumber + 1
	}
	return numberOfOperationsUsingLastNumber + 1
}
