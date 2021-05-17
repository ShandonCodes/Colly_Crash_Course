package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func writeCSV(rows [][]string) {
	f, err := os.Create("./out.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	w := csv.NewWriter(f)

	for _, row := range rows {
		w.Write(row)
	}

	f.Sync()

	w.Flush()

}

func main() {
	c := colly.NewCollector()
	var rows [][]string

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		fmt.Println(string(e.Text))
		rows = append(rows, []string{string(e.Text)})
	})

	c.Visit("https://en.wikipedia.org/wiki/Go_(programming_language)")
	writeCSV(rows)

}
