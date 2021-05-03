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

// 💜 naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	//💜 defer
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
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

	totalLength2, upperName2 := lenAndUpper2("yodi")
	fmt.Println(totalLength2, upperName2)

	repeatMe("yumi", "bread", "icecream", "pasta", "drink")
}
