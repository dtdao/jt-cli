package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(getArticleCmd)
}

var getArticleCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific article by url",
	Long: `This will get you a specific articles.  So instead getting all articles for a particular day,
	This will just get you the specific article given a date`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println(args)
	},
}
