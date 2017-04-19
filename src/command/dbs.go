package command

import (
	"fmt"
	"time"

	"github.com/c88lopez/dbs/src/handlers"
)

// Execute dbs.
func Execute() {
	start := time.Now()

	handlers.Error(Dispatch())

	fmt.Printf("\nElapsed time: %s\n", time.Since(start))
}
