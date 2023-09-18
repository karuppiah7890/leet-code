package twosum

func TwoSum(nums []int32, target int32) []int32 {
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
func TwoSumQuick(nums []int32, target int32) []int32 {
	seenNumbers := map[int32]int32{}

	for index, num := range nums {
		num2 := target - num
		num2Index, ok := seenNumbers[num2]
		if ok {
			return []int32{int32(index), num2Index}
		}
		seenNumbers[num] = int32(index)
	}
	return nil
}

// Brute Force. Time = O(N^2)
func TwoSumBruteForce(nums []int32, target int32) []int32 {
	for index, num := range nums {
		for j := index + 1; j < len(nums); j++ {
			otherNumber := nums[j]
			if num+otherNumber == target {
				return []int32{int32(index), int32(j)}
			}
		}
	}
	return nil
}
