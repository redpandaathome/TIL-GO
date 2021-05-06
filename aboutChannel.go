// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	//go channel
// 	c := make(chan string)
// 	people := [2]string{"yodi", "jordi"}
// 	for _, person := range people {
// 		go isCute(person, c)
// 	}
// 	// result := <-c
// 	// fmt.Println(result)
// 	// receiving a msg...(BLOCKING OPERATION)
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println("waiting for...", i)
// 		fmt.Println(<-c)
// 	}

// }

// func isCute(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + " is cute"
// }
