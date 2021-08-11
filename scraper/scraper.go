package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/schollz/progressbar/v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	japanTimes = "https://www.japantimes.co.jp"
)

type Article struct {
	Title, Content string
}

func ScrapeToday() {
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		if err := os.Mkdir("articles", 0755); err != nil {
			log.Fatal(err)
		}
	}

	bar := progress()

	c := colly.NewCollector(
		colly.AllowedDomains("www.japantimes.co.jp"),
		colly.MaxDepth(2),
	)

	c.OnHTML("li.index-loop-article", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
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
		data := Article{
			Title:   title,
			Content: article,
		}
		jsonFileName := fmt.Sprintf("%s.json", title)
		file, _ := json.MarshalIndent(data, "", "")

		bar.Add(10000)

		//f, err := os.Create(filePath)
		//if err != nil {
		//	panic(err)
		//}

		filePath := fmt.Sprintf("./articles/%s", jsonFileName)
		_ = ioutil.WriteFile(filePath, file, 0644)
		//f.WriteString(article)
	})

	c.Visit(japanTimes)
}

func ScrapeDate(date string) {
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		if err := os.Mkdir("articles", 0755); err != nil {
			panic(err)
		}
	}

	bar := progress()

	c := colly.NewCollector(
		colly.AllowedDomains(),
		colly.MaxDepth(2),
	)
	c.OnHTML("article.story.archive_story", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")

		if strings.Contains(link, date) {
			e.Request.Visit(link)
		}
	})

	//c.OnHTML("article.story.archive_story", func(e *colly.HTMLElement) {
	//	links := e.ChildAttrs("a", "href")
	//	for _, link := range links {
	//		if strings.Contains(link, date) {
	//			e.Request.Visit(link)
	//			fmt.Println(link)
	//		}
	//
	//	}
	//})

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
		bar.Add(10000)

		f, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}

		f.WriteString(article)
	})

	c.Visit(japanTimes + "/news/" + date)
}

func progress() *progressbar.ProgressBar {
	bar := progressbar.DefaultBytes(-1, "Scrapping...")
	return bar
}
