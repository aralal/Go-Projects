package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)
	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText("h2.product-title"))
	})
	c.Visit("https://j2store.net/demo/index.php/shop")

}
