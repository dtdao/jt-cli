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
	Title, Content, Credit, Writer string
}

func ScrapeToday() error {
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		if err := os.Mkdir("articles", 0755); err != nil {
			log.Fatal(err)
		}
	}

	bar := progress()

	c := colly.NewCollector(
		colly.AllowedDomains("japantimes.co.jp", "www.japantimes.co.jp"),
		colly.MaxDepth(1),
	)

	articleColector := c.Clone()

	// Lead stories
	c.OnHTML("div.lead-stories > a.wrapper-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = e.Request.AbsoluteURL(link)
		articleColector.Visit(link)
	})

	// Top stories
	c.OnHTML("div.top-stories > a.wrapper-link.top-story", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = e.Request.AbsoluteURL(link)
		articleColector.Visit(link)
	})

	// Editor picks
	c.OnHTML("div.editors-picks > a.wrapper-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = e.Request.AbsoluteURL(link)
		articleColector.Visit(link)
	})

	/**
	  Japantimes <div.main-content> is separated into different <section>
	  Each section consist of a feature article <div.featured>
	  A subsection list section of articles that relates to that section <ul
	*/
	c.OnHTML("div.featured > > a.wrapper-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = e.Request.AbsoluteURL(link)
		articleColector.Visit(link)
	})

	c.OnHTML("ul.module_articles > li.index-loop-article > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = e.Request.AbsoluteURL(link)
		articleColector.Visit(link)
	})
	/**
	 End of section collector.
	 */

	articleColector.OnHTML("div.main", func(e *colly.HTMLElement) {
		makeArticle(e)
		bar.Add(1)
	})

	return c.Visit(japanTimes)
}

func ScrapeDate(date string) error {
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

	c.OnHTML("div.main", func(e *colly.HTMLElement) {
		makeArticle(e)
		bar.Add(10000)
	})

	return c.Visit(japanTimes + "/news/" + date)
}

func progress() *progressbar.ProgressBar {
	bar := progressbar.DefaultBytes(-1, "Scrapping...")
	return bar
}

func makeArticle(e *colly.HTMLElement) {
	title := e.ChildText("h1")
	credit := e.ChildText("p.credit")
	writer := e.ChildText("h5.writer")
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
		Credit:  credit,
		Writer:  writer,
	}
	jsonFileName := fmt.Sprintf("%s.json", title)
	file, _ := json.MarshalIndent(data, "", "")

	filePath := fmt.Sprintf("./articles/%s", jsonFileName)
	_ = ioutil.WriteFile(filePath, file, 0644)
}
