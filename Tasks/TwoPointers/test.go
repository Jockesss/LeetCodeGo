package main

import (
	"fmt"
)

func main() {
	v := "Hello world"
	for i, j := range v {
		fmt.Println(i, j)
	}
}
