package main

import (
	"encoding/json"
	"fmt" //formatted I/O
	"os"

	"github.com/gocolly/colly" //scraping framework
)

type item struct {
	Name  string `json:name`
	Stars string `json:stars`
	Price string `json:price`
}

func main() {

	c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", r.URL)
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-base", func(_ int, h *colly.HTMLElement) {

			Item := item{
				Name:  h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal"),
				Stars: h.ChildText("span.a-icon-alt"),
				Price: h.ChildText("span.a-price-whole"),
			}
			var Items []item
			Items = append(Items, Item)
			j, err := json.MarshalIndent(Items, " ", " ")
			if err != nil {
				fmt.Println(err)
			}
			os.WriteFile("product.json", j, 0644)
			fmt.Println(string(j))

		})
	})

	c.Visit("https://www.amazon.in/s?k=keyboard")
}
