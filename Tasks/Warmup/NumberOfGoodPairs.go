/*
	Учитывая массив целых чисел nums, верните количество хороших пар.

    Пара (i, j) называется хорошей, если nums[i] == nums[j] и i < j.
*/

package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 1, 1, 3}

	fmt.Println(numIdenticalPairs(n))
}

func numIdenticalPairs(nums []int) int {
	count := 0
	res := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res[nums[i]] = nums[j]
				count++
			}
		}
	}
	return count
}
