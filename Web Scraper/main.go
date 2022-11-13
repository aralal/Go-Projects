package main

import "github.com/gocolly/colly"

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("amazon.in"),
	)
	c.OnHTML(" ", func(h *colly.HTMLElement) {

	})

}
