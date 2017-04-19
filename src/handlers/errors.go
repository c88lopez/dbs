package handlers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Error func
func Error(e error) {
	if nil != e {
		color.Red("Error!")
		fmt.Printf("\n%s\n", e)

		os.Exit(1)
	}
}
