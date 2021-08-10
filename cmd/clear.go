package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(clearCmd)
}



var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear out all scrapped articles",
	Long: `This will clear out all the articles you have scrapped.  
		This is a perminanent action and it cannot be recovered unless you go 
		through and scrape the articles again.`,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := os.Open(Articles)
		if err != nil {
			cmd.ErrOrStderr()
		}

		defer d.Close()
		names, err := d.Readdirnames(-1)
		if err != nil {
			cmd.ErrOrStderr()
		}

		for _, name := range names {
			err = os.RemoveAll(filepath.Join(Articles, name))
			if err != nil {
				cmd.ErrOrStderr()
			}
		}
	},
}
