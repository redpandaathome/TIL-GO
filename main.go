package main

import (
	"fmt"
	"time"
)

func main() {
	// to make them work concurrently...? put 'go'
	// but with both of them with go -> main just finishes
	// main doesn't wait for go routines
	// then how do we let them communicate? use go pipeline!
	go sexyCount("yodi")
	sexyCount("jordi")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy")
		time.Sleep(time.Second)
	}
}
