/*
	Структура slice включает в себя
	type slice struct {
		array unsafe.Pointer  --> Указатель на массив, содержащий элементы slice
		len int --> Длина slice (количество элементов)
		cap int --> Емкость slice
	}
*/

package main

import "fmt"

func main() {
	task1()
	task2()
}

func task1() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[:]
	sl[1] = 10

	fmt.Println(sl)
	fmt.Println(arr, "____________________")
}

func task2() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:]
	sl[1] = 10

	fmt.Println(sl)
	fmt.Println(arr, "____________________")
}

func task3() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:2]
	sl = append(sl, 4)

	fmt.Println(sl)
	fmt.Println(arr, "____________________")
}
