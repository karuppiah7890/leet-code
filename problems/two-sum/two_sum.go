package twosum

func TwoSum(nums []int, target int) []int {
	return TwoSumQuick(nums, target)
}

// Cases to consider
// 1. There can be more than one instance / occurance of the same number
// 2. The target can be 0
// 3. Any number in the array, can be greater than the target,
// or less than the target or even equal to the target. So, any number seen
// has to be recorded to be able to check back for seen numbers when looking
// for a matching number for sum to be equal to target

// Questions to consider
// 1. Can the numbers be negative? Yes! -10^9 <= nums[i] <= 10^9
// 2. Can the target be negative? Yes! -10^9 <= target <= 10^9
// 3. What's the max number of numbers that can be present in the array? 2 <= nums.length <= 10^4

// Quick. Time = O(N)
func TwoSumQuick(nums []int, target int) []int {
	seenSmallNumbers := map[int][]int{}

	for index, num := range nums {
		num2 := target - num
		num2Indices, ok := seenSmallNumbers[num2]
		if ok {
			return []int{index, num2Indices[0]}
		}
		seenSmallNumbers[num] = append(seenSmallNumbers[num], index)
	}
	return nil
}

// Brute Force. Time = O(N^2)
func TwoSumBruteForce(nums []int, target int) []int {
	for index, num := range nums {
		for j := index + 1; j < len(nums); j++ {
			otherNumber := nums[j]
			if num+otherNumber == target {
				return []int{index, j}
			}
		}
	}
	return nil
}
