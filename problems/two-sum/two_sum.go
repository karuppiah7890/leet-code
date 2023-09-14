package twosum

func TwoSum(nums []int, target int) []int {
	return TwoSumV1(nums, target)
}

// Brute Force. O(N^2)
func TwoSumV1(nums []int, target int) []int {
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
