package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/redpandaathome/learngo/scrapper"
)

const fileName string = "jobs.csv"

// Handler
func handleHome(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	// scrapper.Scrape("golang")
	e := echo.New()

	// Routes
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
