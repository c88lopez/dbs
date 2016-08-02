package main

import (
	"io/ioutil"
	"os"

	"fmt"

	"github.com/fatih/color"
)

var mainFolderPath string

type DbsFolder struct {
}

func setMainFolderPath() {
	dir, err := os.Getwd()
	check(err)

	mainFolderPath = dir + "/.dbs"
}

func generateMainFolder() {
	fmt.Print("Initializing... ")

	setMainFolderPath()
	_, err := os.Stat(mainFolderPath)
	if os.IsNotExist(err) {
		newDirPaths := [1]string{"/states"}

		for _, newDirPath := range newDirPaths {
			os.MkdirAll(mainFolderPath+newDirPath, 0775)
		}

		check(ioutil.WriteFile(mainFolderPath+"/config", getConfigTemplate(), 0600))
		check(ioutil.WriteFile(mainFolderPath+"/states/history", []byte{}, 0644))

		color.Green("Done.\n")
	} else {
		color.Yellow("Already initialized!.\n")
	}
}

func getStatesDirPath() string {
	setMainFolderPath()
	return mainFolderPath + "/states"
}
