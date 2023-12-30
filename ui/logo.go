package ui

import "github.com/common-nighthawk/go-figure"

func PrintLogo() {
	myFigure := figure.NewFigure("GeminiShell", "doom", true)

	print("\033[2J")
	print("\033[H")
	print("\n\033[35m")
	myFigure.Print()
	print("\033[0m\n")
}
