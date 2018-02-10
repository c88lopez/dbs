package errors

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Handle func
func Handle(e error) {
	if nil != e {
		color.Red("Error!")
		fmt.Printf("\n%s\n", e)

		os.Exit(1)
	}
}
