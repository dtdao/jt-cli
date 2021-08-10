package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"jt-cli/scraper"

)

func init(){
	rootCmd.AddCommand(todayCmd)
}

var todayCmd = &cobra.Command{
	Use: "today",
	Short: "Get all scrape all articles today",
	Long: "This will scrape japantimes.jp for all articles released today.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This will scrape everything for today")
		scraper.ScrapeToday()
	},
}