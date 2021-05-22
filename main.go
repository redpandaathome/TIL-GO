package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// go get github.com/labstack/echo/v4

// Handler
func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// scrapper.Scrape("golang")
	e := echo.New()

	// Routes
	e.GET("/", handleHome)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
