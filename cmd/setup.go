/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ara-o/hound/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

//TODO: Change project id to project name

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up hound with the proper environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		pineconeHasBeenSetup, _ := pterm.DefaultInteractiveConfirm.Show("Have you created a Pinecone account?")

		if !pineconeHasBeenSetup {
			utils.HandlePineconeSetupInstructions()
			return
		}

		openaiHasBeenSetup, _ := pterm.DefaultInteractiveConfirm.Show("Have you created an Openai account?")

		if !openaiHasBeenSetup {
			utils.HandleOpenaiSetupInstructions()
			return
		}

		pterm.DefaultBox.
			WithRightPadding(10).
			WithLeftPadding(10).
			WithBoxStyle(&pterm.Style{pterm.FgGreen}).
			WithTopPadding(1).
			WithBottomPadding(1).
			Println(pterm.Green("Setup Instructions"))
		pterm.Println("Note: Your information will never be shared and is only stored locally")
		pterm.Println()
		pinecone_project_id, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone project ID")
		pinecone_environment_name, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone environment name")
		pinecone_index_name, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone index name")
		pinecone_api_key, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone api key")
		openai_api_key, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Openai api key")
		pterm.Println(pinecone_api_key, pinecone_environment_name, pinecone_project_id, pinecone_index_name, openai_api_key)

	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
