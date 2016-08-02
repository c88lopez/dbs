package main

import "fmt"

func showHelp() {
	fmt.Print(`Usage: dbs <command> [<arg>, ...]

Commands:

init		Initialize.
config		Configure the database connection.
generate	Build and persist the current database state.
`)
}
