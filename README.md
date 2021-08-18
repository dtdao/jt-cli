# Japan Times Scrapper CLI tool (jt-cli)

### A simple GoLang web scrapping tool for scrapping the japantimes.co.jp website for articles.

Why would I want to build such a tool?

1. Japantimes has a paywall.  A very weak paywall.  While the "paywall" can be disabled
with simple html/scss manipulation.  It does become a chore to do in a browser developer tools.

2. This project is to build on a chrome/firefox extension that ive worked on previously.
That extension does work by manipulating the html/scss and removing the "paywall".

3. I want to continue to learn and build more tools/apps using GoLang.

## Features
- Fast webscrapping using the `Colly` webscrapping framework
- Simple and intuitive commands
- Absolute free!! You never have to pay for any more articles form JapanTimes.
- All articles scraped will be save locally in an `articles` folder, where you will be able to go back and read through  
at your own pace.

```bigquery
Usage:
  jt-cli [command]

Available Commands:
  articles    List out all the articles
  clear       Clear out all scrapped articles
  date        Scrape Japan Times at a particular date
  help        Help about any command
  reader      Turn on reader mode
  today       Get all scrape all articles today
  version     Print the version number of JT-cli tool
```

### Prerequisite 

You might have to install Go on your computer and build  

`go build`



