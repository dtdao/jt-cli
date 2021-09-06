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
	Title, Content, Credit, Writer, Url, Date string
}

func init() {
	if err := makeFile(); err != nil {
		log.Fatal(err)
	}
}

func makeFile() error {
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		if err := os.Mkdir("articles", 0755); err != nil {
			return err
		}
	}
	return nil
}

func makeCollector() (collector *colly.Collector) {
	c := colly.NewCollector(
		colly.AllowedDomains("japantimes.co.jp", "www.japantimes.co.jp"),
		colly.MaxDepth(1),
	)
	return c
}

func ScrapeToday() error {
	bar := progress()

	c := makeCollector()

	articleColector := c.Clone()

	// Lead stories
	c.OnHTML("div.lead-stories > a.wrapper-link", func(e *colly.HTMLElement) {
		visitContent(articleColector, e)
	})

	// Top stories
	c.OnHTML("div.top-stories > a.wrapper-link.top-story", func(e *colly.HTMLElement) {
		visitContent(articleColector, e)
	})

	// Editor picks
	c.OnHTML("div.editors-picks > a.wrapper-link", func(e *colly.HTMLElement) {
		visitContent(articleColector, e)
	})

	/**
	  Japantimes <div.main-content> is separated into different <section>
	  Each section consist of a feature article <div.featured>
	  A subsection list section of articles that relates to that section <ul
	*/
	c.OnHTML("div.featured > > a.wrapper-link", func(e *colly.HTMLElement) {
		visitContent(articleColector, e)

	})

	c.OnHTML("ul.module_articles > li.index-loop-article > a", func(e *colly.HTMLElement) {
		visitContent(articleColector, e)
	})
	/**
	End of section collector.
	*/

	articleColector.OnHTML("div.main", func(e *colly.HTMLElement) {
		makeArticle(e)
		bar.Add(1)
	})

	onRequest(articleColector, bar)

	return c.Visit(japanTimes)
}

func ScrapeUrl(url string) error {
	bar := progress()
	articleCollector := makeCollector()

	articleCollector.OnHTML("div.main", func(e *colly.HTMLElement) {
		makeArticle(e)
	})

	onRequest(articleCollector, bar)
	if err := articleCollector.Visit(url); err != nil {
		return err
	}
	return nil
}
func ScrapeDate(date string) error {
	bar := progress()

	c := makeCollector()

	articleCollector := c.Clone()

	c.OnHTML("article.story.archive_story > div.content_col > header > hgroup > h1 > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, date) {
			articleCollector.Visit(link)
		}
	})

	articleCollector.OnHTML("div.main", func(e *colly.HTMLElement) {
		makeArticle(e)
	})

	onRequest(articleCollector, bar)

	return c.Visit(japanTimes + "/news/" + date)
}

func progress() *progressbar.ProgressBar {
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Scrapping"),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSpinnerType(35),
		progressbar.OptionClearOnFinish())
	return bar
}

func onRequest(c *colly.Collector, bar *progressbar.ProgressBar) {
	c.OnRequest(func(r *colly.Request) {
		message := fmt.Sprintf("Scrapping..%s", r.URL.String())
		bar.Add(1)
		bar.Describe(message)
	})
}

func visitContent(collector *colly.Collector, e *colly.HTMLElement) {
	link := e.Attr("href")
	link = e.Request.AbsoluteURL(link)
	collector.Visit(link)

}

func makeArticle(e *colly.HTMLElement) {
	title := e.ChildText("h1")
	credit := e.ChildText("p.credit")
	writer := e.ChildText("h5.writer")
	date := e.ChildAttr("time", "datetime")
	url := fmt.Sprintf("%s%s", e.Request.URL.Host, e.Request.URL.Path)

	r, size := utf8.DecodeLastRuneInString(title)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	if title[len(title)-1:] == "." {
		title = title[:len(title)-size]
	}

	content := e.ChildText("#jtarticle > p")
	data := Article{
		Title:   title,
		Content: content,
		Credit:  credit,
		Writer:  writer,
		Url:     url,
		Date:    date,
	}
	jsonFileName := fmt.Sprintf("%s.json", title)
	fileExist, err := doesFileExist(jsonFileName)
	if err != nil {
		log.Fatal(err)
	}
	if !fileExist {
		file, _ := json.MarshalIndent(data, "", "")

		filePath := fmt.Sprintf("./articles/%s", jsonFileName)
		_ = ioutil.WriteFile(filePath, file, 0644)
	}

}

func doesFileExist(file string) (bool, error) {
	filePath := fmt.Sprintf("./articles/%s", file)
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	}
	return false, nil
}
