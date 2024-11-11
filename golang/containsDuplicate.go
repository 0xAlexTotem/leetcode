package main

import "sort"

func containsDuplicate(nums []int) bool {

	// Specific case
	if len(nums) <= 1 {
		return false
	}

	// Sorting the slyce
	sort.Ints(nums)

	// Check if there are any duplicate
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
