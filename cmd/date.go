package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jt-cli/scraper"
	"log"
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
			err := scraper.ScrapeToday()
			if err != nil {
				log.Fatal(err)
			}
		}
		var date = args[0]
		err := scraper.ScrapeDate(date)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}
