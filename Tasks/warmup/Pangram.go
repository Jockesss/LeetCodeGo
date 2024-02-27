/*
	A pangram is a sentence where every letter of the English alphabet appears at least once.
	Given a string sentence containing only lowercase English letters,
	return true if sentence is a pangram, or false otherwise.

	Панграмма - это предложение, в котором каждая буква английского алфавита встречается хотя бы один раз.
	Если дано строковое предложение, содержащее только строчные английские буквы, верните true,
	если предложение является панграммой, или false в противном случае.
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var sentence string
	sentence = "leetcode"

	fmt.Println(checkIfPangram(sentence))

}

func checkIfPangram(sentence string) bool {
	seen := make(map[rune]bool)

	for _, num := range strings.ToLower(sentence) {
		if unicode.IsLower(num) {
			seen[num] = true
		}
	}
	return len(seen) == 26
}
