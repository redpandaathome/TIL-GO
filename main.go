package main

import (
	"fmt"
	"time"
)

func main() {
	//go channel
	c := make(chan bool)
	people := [2]string{"yodi", "jordi"}
	for _, person := range people {
		go isCute(person, c)
	}
	// result := <-c
	// fmt.Println(result)
	fmt.Println(<-c)
	fmt.Println(<-c)

}

func isCute(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
