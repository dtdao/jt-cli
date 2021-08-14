package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jt-cli/scraper"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(dateCmd)
}

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Scrape Japan Times at a particular date",
	Long: `Be more specific with you scrapping.  You can provide a date
    value in YYYY/MM/DD format and get all articles for that date.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("Scrapping today")
			return  scraper.ScrapeToday()
		}
		date, err := time.Parse("2006/01/02", args[0])
		if err != nil {
			log.Fatal("Invalid date")
		}
		formattedTime := date.Format("2006/01/02")
		return scraper.ScrapeDate(formattedTime)
	},
}
