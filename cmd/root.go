package cmd

import (
	"log"
	"os"

	"github.com/ara-o/doc-find/utils"
	"github.com/spf13/cobra"
)

var url string
var comprehensive bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doc-find",
	Short: "Easily query through long pages of documentation!",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := utils.ParseURL(url, comprehensive)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Set up flags
func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "Enter the URL of the documentation page you want to query")
	rootCmd.Flags().BoolVarP(&comprehensive, "comprehensive", "c", false, "Select true if you want to search through the entire site's documentation")
	rootCmd.MarkFlagRequired("url")
}
