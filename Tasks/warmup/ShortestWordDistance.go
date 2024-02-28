package main

import (
	"fmt"
	"math"
)

func main() {
	var words = []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	word1 := "fox"
	word2 := "dog"

	fmt.Println(shortestDistance(words, word1, word2))

}

func shortestDistance(arr []string, s1 string, s2 string) int {
	word1 := -1
	word2 := -1
	minDist := math.MaxInt

	for i := 0; i < len(arr); i++ {
		if s1 == arr[i] {
			word1 = i
		} else if s2 == arr[i] {
			word2 = i
		}

		if word1 != -1 && word2 != -1 {
			dist := abs(word2 - word1)
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
