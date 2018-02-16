package main

import (
	"fmt"
	"time"

	"github.com/c88lopez/dbs/cmd"
)

/**
This is only the entry point of the application.
The idea is to handle all the input on a separated component.
*/
func main() {
	start := time.Now()

	cmd.Execute()

	fmt.Printf("\nElapsed time: %s\n", time.Since(start))
}
