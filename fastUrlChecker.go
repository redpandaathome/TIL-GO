// package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errRequestFailed = errors.New("Request failed...")

// type requestResult struct {
// 	url    string
// 	status string
// }

// func main() {
// 	results := make(map[string]string)
// 	c := make(chan requestResult)
// 	// slice...
// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://www.soundcloud.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com/",
// 		"https://nomadcoders.co/",
// 	}

// 	for _, url := range urls {
// 		go hitURL(url, c)
// 	}

// 	for i := 0; i < len(urls); i++ {
// 		result := <-c
// 		results[result.url] = result.status
// 	}

// 	for url, status := range results {
// 		fmt.Println(url, status)
// 	}

// }

// func hitURL(url string, c chan<- requestResult) { //send only
// 	resp, err := http.Get(url)
// 	status := "OK"
// 	if err != nil || resp.StatusCode >= 400 {
// 		status = "FAIL"
// 	}
// 	c <- requestResult{url, status}
// }
