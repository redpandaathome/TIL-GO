package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println(multiply(2, 3))
	// totalLength, _ := lenAndUpper("yumi")
	totalLength, upperName := lenAndUpper("yumi")
	fmt.Println(totalLength, upperName)

	repeatMe("yumi", "bread", "icecream", "pasta", "drink")
}
