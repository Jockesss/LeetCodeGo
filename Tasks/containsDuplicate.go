package main

import (
	"fmt"
)

func main() {
	var test Solution
	arr := []int{1, 2, 3, 3, 4, 5}

	test.containsDuplicate(arr)
	fmt.Println(test.containsDuplicate(arr))
}

type Solution struct{}

// containsDuplicate checks for duplicates in a slice of integers
func (s *Solution) containsDuplicate(nums []int) bool {
	var seen map[int]bool

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = true
	}
	return false
}

//// containsDuplicate checks for duplicates in a slice of integers
//func (s *Solution) containsDuplicate(nums []int) bool {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i] == nums[j] {
//				return true
//			}
//		}
//	}
//	return false
//}
