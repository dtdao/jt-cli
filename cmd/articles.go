package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var Search string

func init() {
	rootCmd.AddCommand(articlesCmd)
	articlesCmd.Flags().StringVarP(&Search, "search", "s", "", "Search for an article with specific term")
}

var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "List out all the articles",
	Long: `All the articles that you have scrapped will be in a local "articles" folder
	This will list out all the articles that you have scrapped.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(Articles)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
	Args: cobra.NoArgs,
}
