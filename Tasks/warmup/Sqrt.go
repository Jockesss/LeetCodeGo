/*
	Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
	The returned integer should be non-negative as well.

	You must not use any built-in exponent function or operator.

	Если задано неотрицательное целое число x, верните квадратный корень из x, округленный до ближайшего целого числа.
	Возвращаемое целое число также должно быть неотрицательным.

	Вы не должны использовать встроенные функции или операторы экспоненты.
*/

package main

import "fmt"

func main() {
	x := 8 // Пример значения
	fmt.Println("Приближенный квадратный корень из", x, ":", mySqrt(x))
}

func mySqrt(x int) int {
	if x < 2 {
		return x // 0 и 1 являются тривиальными случаями
	}

	// Начальные границы поиска
	left, right := 1, x/2

	for left <= right {
		mid := left + (right-left)/2
		sqrt := mid * mid

		if sqrt == x {
			return mid
		} else if sqrt < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
