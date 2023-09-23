package min_operations

func MinOperations(nums []int, x int) int {
	// array of size 1
	if len(nums) == 1 {
		if nums[0] == x {
			// 1 operation to subtract the single number from x
			return 1
		} else {
			// here (nums[0] > x) or (nums[0] < x).
			// hence no operations can be subtracted
			// from x to make it 0
			return -1
		}
	}
	if len(nums) > 1 {
		if nums[0] == x || nums[len(nums)-1] == x {
			// the first or last number could be the same as x.
			// or both the first and last number could be the same as x.
			// In both cases, just 1 operation to subtract the single number from x.
			return 1
		}
		if nums[0] > x && nums[len(nums)-1] > x {
			return -1
		}
	}

	// remove leftmost number and follow that path and try to make x as 0

	// remove rightmost number and follow that path and try to make x as 0

	// if any one - leftmost or rightmost number and
	// more numbers can be subtracted from x to make it 0
	// then return the number of operations.
	// ELSE
	// if both leftmost and rightmost number path does not
	// lead x to become 0 on subtracting numbers, then
	// return -1

	return 0
}
