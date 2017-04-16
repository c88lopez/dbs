package command

import (
	"fmt"
	"time"

	"github.com/c88lopez/dbs/src/handlers"
)

// Executes the dbs.
func Execute() error {
	var err error

	start := time.Now()

	if err = Dispatch(); nil != err {
		handlers.Error(err)
	}

	fmt.Printf("\nElapsed time: %s\n", time.Since(start))

	return nil
}
