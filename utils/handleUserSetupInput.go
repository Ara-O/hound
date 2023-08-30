package utils

import "github.com/pterm/pterm"

type SetupVariables struct {
	Pinecone_project_id       string
	Pinecone_environment_name string
	Pinecone_index_name       string
	Pinecone_api_key          string
	Openai_api_key            string
}

func HandleUserSetupInput() (variables SetupVariables) {
	pterm.DefaultBox.
		WithRightPadding(10).
		WithLeftPadding(10).
		WithBoxStyle(&pterm.Style{pterm.FgGreen}).
		WithTopPadding(1).
		WithBottomPadding(1).
		Println(pterm.Green("Setup Instructions"))
	pterm.Println("Note: Your information will never be shared and will only be stored locally")
	pterm.Println()

	variables = SetupVariables{}
	variables.Pinecone_project_id, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone project ID")
	variables.Pinecone_environment_name, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone environment name")
	variables.Pinecone_index_name, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone index name")
	variables.Pinecone_api_key, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Pinecone api key")
	variables.Openai_api_key, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter your Openai api key")

	return
}
