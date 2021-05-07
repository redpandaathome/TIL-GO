package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://uk.indeed.com/jobs?q=golang&l=United+Kingdom"

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		fmt.Println(getPage(i))
	}
}

// 2. Each page URL
func getPage(page int) string {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	return pageURL
}

// 1. 전체 페이지수
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}
