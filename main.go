package main

import (
	"github.com/c88lopez/dbs/command"
	"github.com/c88lopez/dbs/handlers"
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
