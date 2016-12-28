package handlers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Error(e error) {
	color.Red("Error.")
	fmt.Printf("\n%s\n", e)

	os.Exit(1)
}
