package help

import "fmt"

const Version = "~@DBS_VERSION@~"

// Displays the version of dbs.
func showVersion() {
	fmt.Printf("Version: %s\n", Version)
}

func ShowHelp() {
	showVersion()

	fmt.Print(`
Usage: dbs <command> [<arg>, ...]

Commands:

init		Initialize.
config		Configure the database connection.
generate	Build and persist the current database state.
`)
}
