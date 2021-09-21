package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jt-cli/scraper"
	"log"
	url2 "net/url"
)

var (
	japan_times = "www.japantimes.co.jp"
)

func init() {
	rootCmd.AddCommand(getArticleCmd)
}

var getArticleCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific article by url",
	Long: `This will get you a specific articles.  So instead getting all articles for a particular day,
	This will just get you the specific article given a date`,
	Run: func(cmd *cobra.Command, args []string) {
		for i, url := range args {
			url, err := url2.ParseRequestURI(url)
			if err != nil {
				fmt.Printf("%s, is not a valid url. \n", args[i])
				continue
			}
			if url.Host != japan_times {
				log.Fatal("You did not provide a valid JapanTimes url")
			}
			articleUrl := fmt.Sprintf("https://%s%s", japan_times, url.Path)
			scraper.ScrapeUrl(articleUrl)
		}
	},
}
