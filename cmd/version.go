package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(versionCmd)
}


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of JT-cli tool",
	Long:  `All software has versions. This is jt-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jt-cli version 1.0")
	},
	Args: cobra.NoArgs,
}