package main

import "fmt"

// Ascending

// Loop through your array
// Do the following N times, starting with first element. N is the number of elements
// For each set of consecutive elements, check the bigger value and put it on the right

func main() {
	array := []int{5, 9, 8, 7, 1, 3, 4}
	fmt.Println("Before sorting: ", array)
	bubbleSort(array)
	fmt.Println("After sorting: ", array)
}

// ascending order
func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for index := 0; index < len(array)-1-i; index++ {
			if array[index] > array[index+1] {
				array[index], array[index+1] = array[index+1], array[index]
			}
		}
	}
}
