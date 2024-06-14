package main

import (
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

var c = cache.New(24*time.Hour, 30*time.Minute)

func getRecyclingLocations() ([]byte, error) {
	x, found := c.Get("recyclingLocations")
	if found {
		return x.([]byte), nil
	}

	resp, err := http.Get("https://data.cityofnewyork.us/resource/sxx4-xhzg.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.Set("recyclingLocations", body, 24*time.Hour)

	return body, nil
}

func main() {
	e := echo.New()
	e.GET("/recycling_locations", func(c echo.Context) error {
		body, err := getRecyclingLocations()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, string(body))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
