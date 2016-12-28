package handlers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Error(e error) {
	color.Red("\n\nError")
	fmt.Println(e)

	os.Exit(1)
}
