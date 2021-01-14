package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Stock struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Volume string `json:"Volume"`
}

var ticker string

func main() {
	allStocks := make(map[string]Stock)

	collector := colly.NewCollector(colly.AllowedDomains("cnbc.com", "www.cnbc.com"))

	collector.OnHTML(".symbol", func(element *colly.HTMLElement) {
		ticker = element.Text
		allStocks[ticker] = Stock{}
	})

	collector.OnHTML(".name", func(element *colly.HTMLElement) {
		name := element.Text
		if stock, ok := allStocks[ticker]; ok {
			stock.Name = name
			allStocks[ticker] = stock
		}
	})

	collector.OnHTML(".last original ng-binding", func(element *colly.HTMLElement) {
		fmt.Println("price part is reaching")
		price := element.Text
		if stock, ok := allStocks[ticker]; ok {
			stock.Price = price
			allStocks[ticker] = stock
		}
	})

	collector.OnHTML(".volume original ng-binding", func(element *colly.HTMLElement) {
		volume := element.Text
		if stock, ok := allStocks[ticker]; ok {
			stock.Volume = volume
			allStocks[ticker] = stock
		}
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})
	collector.Visit("https://www.cnbc.com/quotes/?symbol=AAPL")
	fmt.Print(allStocks)
}
