package cmd

import (
	"github.com/spf13/cobra"
	"jt-cli/scraper"
	"strings"
)

func init(){
	rootCmd.AddCommand(dateCmd)
}

var dateCmd = &cobra.Command{
	Use: "date",
	Short: "Scrape Japan Times at a particular date",
	Long:  `Be more specific with you scrapping.  You can provide a date
    value in YYYY/MM/DD format and get all articles for that date.`,
	Run: func(cmd *cobra.Command, args []string) {
		date := strings.Join(args, "")
		//scraper.ScrapeDate(date)
		scraper.ScrapeDate(date)
	},
}