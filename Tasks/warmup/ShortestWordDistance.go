package main

import "fmt"

func main() {
	var words = []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	word1 := "fox"
	word2 := "dog"

	fmt.Println(shortestDistance(words, word1, word2))

}

func shortestDistance(arr []string, s1 string, s2 string) int {
	var res int
	word1 := 0
	word2 := 0

	for i := 0; i < len(arr); i++ {
		if s1 == arr[i] {
			word1 = i
		}
		if s2 == arr[i] {
			word2 = i
		}
	}

	res = word2 - word1
	return res
}
