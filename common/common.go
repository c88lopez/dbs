package common

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c88lopez/dbs/config"
	"github.com/fatih/color"
)

var mainFolderPath string

func SetMainFolderPath() error {
	dir, err := os.Getwd()
	if nil != err {
		return err
	}

	mainFolderPath = dir + "/.dbs"

	return nil
}

func GenerateMainFolder() error {
	fmt.Print("Initializing... ")

	SetMainFolderPath()
	_, err := os.Stat(mainFolderPath)
	if os.IsNotExist(err) {
		newDirPaths := [2]string{"/states", "/logs"}

		for _, newDirPath := range newDirPaths {
			os.MkdirAll(mainFolderPath+newDirPath, 0775)
		}

		template, err := config.GetConfigTemplate()
		if nil != err {
			return err
		}

		err = ioutil.WriteFile(mainFolderPath+"/config", template, 0600)
		if nil != err {
			return err
		}

		err = ioutil.WriteFile(mainFolderPath+"/states/history", []byte{}, 0644)
		if nil != err {
			return err
		}

		color.Green("Done.\n")
	} else {
		color.Yellow("Already initialized!\n")
	}

	return nil
}

func GetStatesDirPath() string {
	SetMainFolderPath()
	return mainFolderPath + "/states"
}
