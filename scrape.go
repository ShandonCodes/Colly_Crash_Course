package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		fmt.Println(string(e.Text))
	})

	c.Visit("hps://en.wikipedia.org/wiki/Go_(programming_language)")

}
