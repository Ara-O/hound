package utils

import "github.com/pterm/pterm"

func HandleOpenaiSetupInstructions() {
	pterm.DefaultBox.
		WithBoxStyle(&pterm.Style{pterm.FgRed}).
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(1).
		WithBottomPadding(1).
		Println(pterm.Red("These steps will be necessary"))
	pterm.Println("[ 1 ] Create your OpenAI account here - https://platform.openai.com/")
	pterm.Println("[ 2 ] Once an account is created, go to the View API keys option under your account")
	pterm.Println("[ 3 ] An API key will be needed to run hound. You will be prompted for a new secret key during the setup step")

	pterm.Println()
	pterm.Println(pterm.LightGreen("Now you can proceed to setup hound 🎉"))
}
