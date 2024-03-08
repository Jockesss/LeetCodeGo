/*
	Given an array of sorted numbers, move all non-duplicate number instances at the beginning of the array in-place.
	The relative order of the elements should be kept the same and you should not use any extra space so
	that the solution has constant space complexity i.e., .

	Move all the unique number instances at the beginning of the array and
	after moving return the length of the subarray that has no duplicate in it.

	Учитывая массив отсортированных чисел, переместите все недублирующиеся экземпляры чисел в начало массива на место.
	Относительный порядок элементов должен оставаться неизменным, и вы не должны использовать лишнее пространство,
	чтобы решение имело постоянную пространственную сложность, т.е. .

	Переместите все уникальные экземпляры чисел в начало массива и
	после перемещения верните длину подмассива, в котором нет дубликатов.

*/

package main

import "fmt"

func main() {
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	fmt.Println(remove(arr))
}

func remove(arr []int) int {
	m := make(map[int]bool)
	count := 0
	//start := 0
	//end := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			count++
		}
		m[arr[i]] = true
	}

	return count
}
