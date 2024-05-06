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
	task5()
}

func task1() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[:]
	sl[1] = 10

	fmt.Println(sl)
	fmt.Println(arr)
	fmt.Println("____________________")
}

func task2() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:]
	sl[1] = 10

	fmt.Println(sl)
	fmt.Println(arr)
	fmt.Println("____________________")
}

func task3() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:2]
	sl = append(sl, 4)

	fmt.Println(sl)
	fmt.Println(arr)
	fmt.Println("____________________")
}

func task4() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:2]
	sl2 := append(sl, 4)
	sl2[0] = 8
	fmt.Println(arr)
	fmt.Println(sl)
	fmt.Println(sl2)
	fmt.Println("____________________")
}

func task5() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:2]
	sl = append(sl, 9)
	sl = append(sl, 8)
	sl = append(sl, 7)
	fmt.Println(arr)

	sl = append(sl, 6)
	fmt.Println(arr)

	sl[0] = 0
	sl[1] = 0

	fmt.Println(sl)
}
