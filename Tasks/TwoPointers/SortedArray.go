/*
Given a sorted array, create a new array containing
squares of all the numbers of the input array in the sorted order.

Задав отсортированный массив, создайте новый массив,
содержащий квадраты всех чисел входного массива в отсортированном порядке.
*/
package main

import "fmt"

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	sortedArr := makeSquares(arr)
	fmt.Println(sortedArr) // Вывод: [0 1 9 16 100]
}

func makeSquares(arr []int) []int {
	n := len(arr)
	squares := make([]int, n)

	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		if abs(arr[left]) > abs(arr[right]) {
			squares[i] = arr[left] * arr[left]
			left++
		} else {
			squares[i] = arr[right] * arr[right]
			right--
		}
	}

	return squares
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
