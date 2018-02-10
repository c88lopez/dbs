package command

import (
	"fmt"
	"time"

	"github.com/c88lopez/dbs/src/errors"
)

// Execute dbs.
func Execute() {
	start := time.Now()

	errors.Handle(Dispatch())

	fmt.Printf("\nElapsed time: %s\n", time.Since(start))
}
