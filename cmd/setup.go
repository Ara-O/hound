/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

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

		path := utils.GetEnvironmentVariablePath()

		data := utils.HandleUserSetupInput()

		storeEnvData(path, data)

		pterm.Println()
		pterm.Println(pterm.LightGreen("Success! All environment variables were successfully stored in ", path))
		pterm.Println(pterm.LightGreen("You can now run hound 🎉", path))
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

func storeEnvData(path string, data utils.SetupVariables) {

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	envFileContent := fmt.Sprintf(
		`MAX_DEPTH=10
PINECONE_API_KEY=%s
PINECONE_PROJECT_ID=%s
PINECONE_ENVIRONMENT_NAME=%s
PINECONE_INDEX_NAME=%s
OPENAI_API_KEY=%s`,
		data.Pinecone_api_key,
		data.Pinecone_project_id,
		data.Pinecone_environment_name,
		data.Pinecone_index_name,
		data.Openai_api_key,
	)
	file.WriteString(envFileContent)
}
