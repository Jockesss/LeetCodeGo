package main

import "fmt"

func main() {
	//arr := []int{3, 2, 5, 6, 21}
	//target := 11
	//
	//fmt.Println(twoSum(arr, target))
	var result map[int]bool
	l := result[10]
	fmt.Println(l)
}

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return append(result, i, j)
			}
		}
	}
	return result
}

//func twoSum(nums []int, target int) []int {
//	var result []int
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return append(result, i, j)
//			}
//		}
//	}
//	return result
//}
