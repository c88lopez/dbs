package help

import "fmt"

const version = "~@DBS_VERSION@~"

// Displays the version of dbs.
func showVersion() {
	fmt.Printf("Version %s\n", version)
}

// ShowHelp func
func ShowHelp() {
	showVersion()

	fmt.Print(`
Usage: dbs <command> [<arg>, ...]

Commands:

init	Initialize.
config	Configure the database connection.
new		Build and persist the current database state.
`)
}
