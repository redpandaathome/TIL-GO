package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	url      string
	title    string
	location string
	salary   string
	sumamry  string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://uk.indeed.com/jobs?q=" + term + "&l=United+Kingdom"
	startTime := time.Now()
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
		// jobs = append(jobs, c...)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("DONE! ", len(jobs))
	endTime := time.Now()
	fmt.Println("Operation time: ", endTime.Sub(startTime))

	// writing + channel 전 기록 📝
	// 	DONE!  75
	// Operation time:  7.520386795s

	// writing + channel 후 기록 📝
	// DONE!  75
	// Operation time:  5.684296518s
}

// 2. Each page URL
func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	//✨
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*10)
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

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Find("h2>a").Attr("href")
	url := "https://uk.indeed.com" + id
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".location").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary>ul>li").Text())

	// fmt.Println(url, title, location, salary, summary)
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
func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
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

func writeJobs(jobs []extractedJob) {
	//💜
	c := make(chan []string)
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"URL", "Title", "Location", "Salary", "Summary"}
	errWrite := w.Write(headers)
	checkErr(errWrite)
	for _, job := range jobs {
		//✨ go routine으로 개선해보기.
		go writeJobDetail(job, c)
	}

	for i := 0; i < len(jobs); i++ {
		jobData := <-c
		errJobWrite := w.Write(jobData)
		checkErr(errJobWrite)
	}
}

func writeJobDetail(job extractedJob, c chan<- []string) {
	c <- []string{job.url, job.title, job.location, job.salary, job.sumamry}
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
