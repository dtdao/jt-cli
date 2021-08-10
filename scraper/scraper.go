package scraper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"unicode/utf8"
)

func ScrapeToday() {
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		if err := os.Mkdir("articles", 0755); err != nil {
			panic(err)
		}
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.japantimes.co.jp"),
		colly.MaxDepth(2),
	)

	c.OnHTML("li.index-loop-article", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a","href")
		e.Request.Visit(link)
	})

	c.OnHTML("div.main", func(e *colly.HTMLElement) {
		title := e.ChildText("h1")
		r, size := utf8.DecodeLastRuneInString(title)
		if r == utf8.RuneError && (size == 0 || size == 1) {
			size = 0
		}
		if title[len(title)-1:] == "." {
			title = title[:len(title)-size]
		}

		article := e.ChildText(`div[id="jtarticle"] p`)
		fileName := fmt.Sprintf("%s.txt", title)
		filePath := fmt.Sprintf("./articles/%s", fileName)

		f, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}

		f.WriteString(article)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnScraped(func(r *colly.Response) {
		r.Request.Visit("/news")
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.japantimes.co.jp/")

}
