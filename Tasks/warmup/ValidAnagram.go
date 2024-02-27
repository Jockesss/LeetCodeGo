/*
	Given two strings s and t, return true if t is an anagram of s, and false otherwise.

	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
	typically using all the original letters exactly once.

	Если даны две строки s и t, верните true, если t является анаграммой s, и false в противном случае.

	Анаграмма - это слово или фраза, образованная путем перестановки букв другого слова или фразы,
	обычно с использованием всех исходных букв ровно один раз.
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "a"
	t := "ab"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	resS := make(map[rune]int)
	resT := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, sym := range clearString(s) {
		resS[sym]++
	}
	for _, sym := range clearString(t) {
		resT[sym]++
	}

	for k, v1 := range resS {
		if v2, ok := resT[k]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}

func clearString(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Println("Ошибка компиляции регулярного выражения:", err)
	}

	return strings.ToLower(reg.ReplaceAllString(s, ""))
}
