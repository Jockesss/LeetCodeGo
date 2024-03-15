/*
	Given an array of unsorted numbers and a target number,
	find a triplet in the array whose sum is as close to the target number as possible,
	return the sum of the triplet.
	If there are more than one such triplet, return the sum of the triplet with the smallest sum.

	Если задан массив неотсортированных чисел и целевое число, найдите в массиве тройку,
	сумма которой максимально близка к целевому числу, и верните сумму этой тройки.
	Если таких триплетов несколько, верните сумму триплета с наименьшей суммой.

*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-2, 0, 1, 2}
	target := 2
	fmt.Println("Сумма ближайшей тройки:", searchTriplet(nums, target))
}

func searchTriplet(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			currentSum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(target-currentSum)) < math.Abs(float64(target-closestSum)) {
				closestSum = currentSum
			}
			if currentSum < target {
				l++
			} else {
				r--
			}
		}
	}

	return closestSum
}
