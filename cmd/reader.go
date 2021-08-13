package cmd

import (
	"github.com/spf13/cobra"
	"jt-cli/reader"
)


func init(){
	rootCmd.AddCommand(readerCmd)
}

var readerCmd = &cobra.Command{
	Use: "reader",
	Short: "Turn on reader mode",
	Long: "Search, select and read your articles",
	Run: func(cmd *cobra.Command, args []string) {
		reader.Reader()
	},
	Args: cobra.NoArgs,
}