package main

import (
	"encoding/json"
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

func generateInitFolder() {
	fmt.Print("Initializing... ")

	setMainFolderPath()
	if _, err := os.Stat(mainFolderPath); os.IsNotExist(err) {
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

func getConfigTemplate() []byte {
	baseConfig := Config{
		Driver:   "mysql",
		Username: "dummy",
		Password: "dummy",
		Database: "dummy",
	}

	json, err := json.Marshal(baseConfig)
	check(err)

	return json
}

func getStatesDirPath() string {
	setMainFolderPath()
	return mainFolderPath + "/states"
}
