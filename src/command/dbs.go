package command

import (
	"fmt"
	"time"
)

// Executes the dbs.
func Execute() error {
	start := time.Now()

	Dispatch()

	fmt.Printf("\nElapsed time: %s\n", time.Since(start))
	return nil
}
