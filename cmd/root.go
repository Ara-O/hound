package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doc-find",
	Short: "Easily query through long pages of documentation!",
	// Long:  `Tired of struggling through lines of brain-numbing text just to find the exact piece of information you need about a particular subject? Well, tire no more. Easily look through hundreds of lines of documentation through an interactive interface`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sp")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "Enter the URL of the documentation page you want to query")
	rootCmd.Flags().BoolP("comprehensive", "c", false, "Select true if you want to search through the entire site's documentation")
	rootCmd.MarkFlagRequired("url")
}
