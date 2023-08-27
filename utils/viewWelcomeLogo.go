package utils

import "github.com/pterm/pterm"

func ViewWelcomeLogo() {
	pterm.DefaultBox.
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(1).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		Println("Welcome to Doc-Find!")
}
