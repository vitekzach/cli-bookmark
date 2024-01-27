package src

import (
	"fmt"

	"github.com/fatih/color"
)

var w *color.Color

func init() {
	w = color.New(color.FgRed, color.Bold, color.Underline)
}

func printWarning(text string) {
	w.Println(fmt.Sprintf("WARNING: %v", text))
}

func printError(text string) {
	w.Println(fmt.Sprintf("ERROR: %v", text))
}
