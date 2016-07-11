package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"fmt"

	"github.com/fatih/color"
)

var mainFolderPath string

type Version struct {
}

func setMainFolderPath() {
	dir, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	mainFolderPath = dir + "/.dbs"
}

func generateInitFolder() {
	fmt.Printf("Initializing... ")

	setMainFolderPath()
	if _, err := os.Stat(mainFolderPath); os.IsNotExist(err) {
		newDirPaths := [1]string{"/states"}

		for _, newDirPath := range newDirPaths {
			os.MkdirAll(mainFolderPath+newDirPath, 0775)
		}

		newFilePaths := [2]string{"/config", "/states/history"}

		for _, newFilePath := range newFilePaths {
			err = ioutil.WriteFile(mainFolderPath+newFilePath, getConfigTemplate(), 0644)
			if err != nil {
				log.Panic(err)
			}
		}

		color.Green("Done.\n")
	} else {
		color.Yellow("Already initialized!.\n")
	}
}

func getConfigTemplate() []byte {
	baseConfig := Config{
		Driver:   "mysql",
		Username: "root",
		Password: "root",
		Database: "test",
	}

	json, err := json.Marshal(baseConfig)
	if err != nil {
		log.Panic(err)
	}

	return json
}

func getStatesDirPath() string {
	return mainFolderPath+"/states"
}