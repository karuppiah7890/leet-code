package array

func buildArray(nums []int) []int {
	for index := 0; index < len(nums); index++ {
		if nums[index] < 0 {
			// ^^^^ WE MOVE ON AS THIS IS MARKED AS TRAVELLED TO OR
			// VISITED, AND UPDATED OR GOTTEN NEW VALUE
			continue
		}

		if index == nums[index] {
			nums[index] = -nums[index] - 1
			// ^^^^ TO MARK THIS AS TRAVELLED TO OR VISITED, AND UPDATED
			// OR GOTTEN NEW VALUE
			continue
		}

		// DON'T need an else of above condition
		// or an if condition, to check if index is
		// not equal to the value at index.
		// So, if the execution of the code comes
		// here, it means index is not equal to the value
		// at index.

		// Can also be called as startPointIndex and startPointIndexOldValue.
		// Can also be called "loop" or "cycle", instead of "stop", "start" etc
		stopPointIndex := index
		stopPointIndexOldValue := nums[index]

		for {
			nextIndex := nums[index]
			if nums[index] == stopPointIndex {
				nums[index] = stopPointIndexOldValue
				// ^^^^ AS stopPointIndexOldValue is the old value or original
				// value at stopPointIndex and we need the original value
				// of nums[nums[index]] not the overwritten one with a negative
				// sign. This is why we store the old value, so that it can
				// be retrieved later
			} else {
				nums[index] = nums[nums[index]] // AS nums[index] = value
			}

			nums[index] = -nums[index] - 1
			// ^^^^ TO MARK THIS AS TRAVELLED TO OR VISITED, AND UPDATED
			// OR GOTTEN NEW VALUE

			index = nextIndex

			if index == stopPointIndex {
				break // break out of loop
			}
		}
	}

	for index, value := range nums {
		nums[index] = -value - 1
	}

	return nums
}
