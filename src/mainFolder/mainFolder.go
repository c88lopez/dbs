package mainFolder

import (
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

var (
	mainFolderPath, configFilePath string
)

// Generate func
func Generate() error {
	var (
		err      error
		template []byte
	)

	setMainFolderPath()
	_, err = os.Stat(mainFolderPath)
	if os.IsNotExist(err) {
		newDirPaths := [1]string{"/states"}

		for _, newDirPath := range newDirPaths {
			os.MkdirAll(mainFolderPath+newDirPath, 0775)
		}

		template, err = GetConfigTemplate()
		if nil == err {
			if err = ioutil.WriteFile(GetConfigFilePath(), template, 0600); nil == err {
				if err = ioutil.WriteFile(GetHistoryFilePath(), []byte{}, 0644); nil == err {
					color.Green("Done.\n")
				}
			}
		}
	} else {
		color.Yellow("Already initialized!\n")
	}

	return err
}

// GetMainFolderPath returns the absolute path of .dbs
func GetMainFolderPath() string {
	if mainFolderPath == "" {
		setMainFolderPath()
	}

	return mainFolderPath
}

// GetConfigFilePath gets the global `configFilePath`
func GetConfigFilePath() string {
	return GetMainFolderPath() + "/config"
}

// GetStatesDirPath func
func GetStatesDirPath() string {
	return GetMainFolderPath() + "/states"
}

// GetHistoryFilePath func
func GetHistoryFilePath() string {
	return GetStatesDirPath() + "/history"
}

// setMainFolderPath func
func setMainFolderPath() {
	dir, err := os.Getwd()
	if nil != err {
		panic(err)
	}

	mainFolderPath = dir + "/.dbs"
}
