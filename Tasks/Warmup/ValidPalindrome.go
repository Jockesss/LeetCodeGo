/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Фраза является палиндромом, если после преобразования всех заглавных букв в строчные и удаления всех неалфавитно-цифровых символов,
она читается одинаково и вперед, и назад. К буквенно-цифровым символам относятся буквы и цифры.

Если задана строка s, верните true, если она является палиндромом, или false в противном случае.
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "aaa"

	fmt.Println(isPalindrome(s))

}

func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println("Ошибка компиляции регулярного выражения:", err)
		return false
	}

	s = strings.ToLower(reg.ReplaceAllString(s, ""))
	if s == "" {
		return true
	}

	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
