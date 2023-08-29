package utils

import "github.com/pterm/pterm"

func HandlePineconeSetupInstructions() {
	pterm.DefaultBox.
		WithBoxStyle(&pterm.Style{pterm.FgRed}).
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(1).
		WithBottomPadding(1).
		Println(pterm.Red("These steps will be necessary"))
	pterm.Println("[ 1 ] Create your Pinecone account here - https://www.pinecone.io/")
	pterm.Println("[ 2 ] Once an account is created, delete the initial project")
	pterm.Println("[ 3 ] Create a new project")
	pterm.Println("[ 4 ] Once the projects have initialized, ")
	pterm.Println("[ 5 ] Create a new index and choose 1536 for the dimensions")
	pterm.Println("[ 6 ] Once the index has been created, create a new API key")
	pterm.Println()
	pterm.Println(pterm.LightGreen("Now you should have all the Pinecone variables necessary 🎉"))
}
