package main

import (
	"fmt"
	"time"
)

func main() {
	//go channel
	c := make(chan string)
	people := [2]string{"yodi", "jordi"}
	for _, person := range people {
		go isCute(person, c)
	}
	// result := <-c
	// fmt.Println(result)
	// waiting for a msg...(BLOCKING OPERATION)
	resultOne := <-c
	resultTwo := <-c
	fmt.Println("Received...", resultOne)
	fmt.Println("Received...", resultTwo)

}

func isCute(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is cute"
}
