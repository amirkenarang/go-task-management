package utils

import "github.com/fatih/color"

func LogSuccess(msg string, parameter ...string) {

	message := msg
	if len(parameter) > 0 {
		message += " " + parameter[0]
	}

	color.Cyan("┌──────────────────────────────────────┐")
	color.Cyan("     " + message + "     ")
	color.Cyan("└──────────────────────────────────────┘")

}
