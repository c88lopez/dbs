package main

import "fmt"

func showHelp() {
	var helpDescription string
	helpDescription += "Usage: dbs <command> [<arg>, ...]\n"
	helpDescription += "\n"
	helpDescription += "Commands:\n"
	helpDescription += "\n"
	helpDescription += "init\t\t\tInitialize.\n"
	helpDescription += "config\t\t\tConfigure the database connection.\n"
	helpDescription += "build-schema-state\tBuild and persist the current database state.\n"
	helpDescription += "\n"

	fmt.Printf("%s", helpDescription)
}
