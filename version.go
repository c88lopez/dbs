package main

import (
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
		newDirPaths := [2]string{"/states", "/history"}

		for _, newDirPath := range newDirPaths {
			os.MkdirAll(mainFolderPath+newDirPath, 0775)
		}

		newFilePaths := [1]string{"/config"}

		for _, newFilePath := range newFilePaths {
			err = ioutil.WriteFile(mainFolderPath+newFilePath, []byte{}, 0775)
			if err != nil {
				log.Panic(err)
			}
		}

		color.Green("Done.\n")
	} else {
		color.Yellow("Already initialized!.\n")
	}
}
