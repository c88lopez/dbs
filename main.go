package dbs

import (
	"github.com/c88lopez/dbs/src/command"
	"github.com/c88lopez/dbs/src/handlers"
)

/**
This is only the entry point of the application.
The idea is to handle all the input on a separated component.
*/
func main() {
	if err := command.Execute(); nil != err {
		handlers.Error(err)
	}
}
