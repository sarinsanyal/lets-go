// slices
package main

import (
	"fmt"
	// "slices"
)

func main() {
	fmt.Println("Function starts here: ")
	manual_slices()
	slicing_existing_array()
	appending_slice()
	make_slice()
}

func manual_slices() {
	// manual slices
	var slc1 []int = []int{1, 2, 3}
	fmt.Println("\nSlice is:", slc1)
}

func slicing_existing_array() {
	//slicing an existing array
	numbers := [5]string{"A", "B", "C", "D", "E"}
	subSlice := numbers[1:4]

	fmt.Println("\nArray of Numbers is:", numbers)
	fmt.Println("Slice of Numbers from array is:", subSlice)

	subSlice[0] = "Z"

	fmt.Println("Updated Slice of Numbers from array is:", subSlice)
	fmt.Println("Updated Array of Numbers is:", numbers)
}

func appending_slice() {
	var items []string

	items = append(items, "Apple")
	fmt.Println("\nAppended Slice is:", items)
	items = append(items, "Banana", "Sarin")
	fmt.Println("Newly Appended Slice is:", items)
}

func make_slice() {
	nums := make([]int, 3, 5) // make([]Type, length, capacity)
	nums[0] = 100
	nums[1] = 200
	nums[2] = 300
	fmt.Printf("\nSlice: %v | Length: %d | Capacity: %d", nums, len(nums), cap(nums))
}
