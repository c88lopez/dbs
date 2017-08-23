package help

import "fmt"

const version = "b93e2dd4cc2942e44a4cfdbedb368e419d432d1b"

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

init      Initialize.
config    Configure the database connection.
new       Build and persist the current database state.
`)
}
