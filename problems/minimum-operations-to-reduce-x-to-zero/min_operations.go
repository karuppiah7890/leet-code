package min_operations

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

	return minOperations(nums, x, 0, len(nums)-1)
}

// the start index and end index are both inclusive indices
func minOperations(nums []int, x int, startIndex int, endIndex int) int {
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

	// the first and/ the last number is less than x

	// remove first number and follow that path and try to make x as 0
	numberOfOperationsUsingFirstNumber := minOperations(nums, x-first, startIndex+1, endIndex)

	// remove last number and follow that path and try to make x as 0
	numberOfOperationsUsingLastNumber := minOperations(nums, x-last, startIndex, endIndex-1)

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
