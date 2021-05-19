package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	url      string
	title    string
	location string
	salary   string
	sumamry  string
}

var baseURL string = "https://uk.indeed.com/jobs?q=golang&l=United+Kingdom"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		// for i := 0; i < 1; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	// fmt.Println(jobs)
	writeJobs(jobs)
	fmt.Println("DONE! ", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	//💜
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"URL", "Title", "Location", "Salary", "Summary"}
	errWrite := w.Write(headers)
	checkErr(errWrite)
	for _, job := range jobs {

		jobSlice := []string{job.url, job.title, job.location, job.salary, job.sumamry}
		errJobWrite := w.Write(jobSlice)
		checkErr(errJobWrite)
	}
}

// 2. Each page URL
func getPage(page int) []extractedJob {
	var jobs []extractedJob
	//✨
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println("pageURL:", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	fmt.Println("searchCards...", searchCards)
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	return jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Find("h2>a").Attr("href")
	url := "https://uk.indeed.com" + id
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".location").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary>ul>li").Text())

	fmt.Println(url, title, location, salary, summary)
	// return extractedJob{
	// ✨
	c <- extractedJob{
		url,
		title,
		location,
		salary,
		summary}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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
