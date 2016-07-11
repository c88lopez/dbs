package main

import "fmt"

func showHelp() {
	var helpDescription string
	helpDescription += "Usage: dbs <command> [<arg>, ...]\n"
	helpDescription += "\n"
	helpDescription += "Commands:\n"
	helpDescription += "\n"
	helpDescription += "init\t\t\tInitialize.\n"
	helpDescription += "generate-config-file\tCreates the config file.\n"
	helpDescription += "\n"

	fmt.Printf("%s", helpDescription)
}
