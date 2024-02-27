/*
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Получив строку s, переверните только все гласные в строке и верните ее.

Гласные - это 'a', 'e', 'i', 'o' и 'u', и они могут встречаться как в нижнем, так и в верхнем регистре, причем более одного раза.
*/
package main

import (
	"fmt"
)

func main() {
	s := "aA"
	fmt.Println(reverseVowels(s))
}

func reverseVowels(s string) string {
	var r []rune
	r = []rune(s)
	start := 0
	end := len(r) - 1

	for start < end {
		if vowels(r[start]) && vowels(r[end]) {
			r[start], r[end] = r[end], r[start]
			start++
			end--
		} else {
			if !vowels(r[start]) {
				start++
			}
			if !vowels(r[end]) {
				end--
			}
		}
	}

	return string(r)
}

func vowels(r rune) bool {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	_, ok := vowels[r]
	return ok
}
