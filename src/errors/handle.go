package errors

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	// ErrNotInitialized error
	ErrNotInitialized = errors.New("Dbs was not initialized")
)

// Handle func
func Handle(err error) {
	if err != nil {
		color.Red("Error!")

		switch err {
		case ErrNotInitialized:
			fmt.Printf("\n%s\n", err)
			break

		default:
			fmt.Printf("\nUnhandled / Unexpected error: %s\n", err)
		}

		os.Exit(1)
	}
}
