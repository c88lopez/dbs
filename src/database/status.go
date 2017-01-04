package database

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"

	"github.com/c88lopez/dbs/src/common"
	"github.com/c88lopez/dbs/src/config"
	"github.com/c88lopez/dbs/src/entity"
	"github.com/c88lopez/dbs/src/handlers"

	"github.com/fatih/color"
)

func New() error {
	var err error

	config.SetConfigFilePath()
	loadConfiguration()

	if err := StartConnectionPool(); nil != err {
		return err
	}
	defer CloseConnectionPool()

	s, err := BuildSchemaState()
	if nil != err {
		return err
	}

	err = generateJsonSchemaState(s)

	return err
}

func generateJsonSchemaState(s *entity.Schema) error {
	fmt.Print("Generating json and json hash... ")

	var err error

	schemaJson, err := json.Marshal(s)
	if nil != err {
		return err
	}

	jsonHash := sha1.Sum(schemaJson)

	statesDirPath := common.GetStatesDirPath()
	jsonFilePath := fmt.Sprintf("%v%c%x", statesDirPath, os.PathSeparator, jsonHash)

	historyFilePath := fmt.Sprintf("%v%c%v", statesDirPath, os.PathSeparator, "history")
	if _, err = os.Stat(jsonFilePath); os.IsNotExist(err) {
		jsonFile, err := os.Create(jsonFilePath)
		if nil != err {
			return err
		}
		defer jsonFile.Close()

		_, err = jsonFile.Write(schemaJson)
		if nil != err {
			return err
		}

		historyFile, err := os.OpenFile(historyFilePath, os.O_WRONLY|os.O_APPEND, 0644)
		if nil != err {
			return err
		}
		defer historyFile.Close()

		historyFile.WriteString(fmt.Sprintf("%x\n", jsonHash))

		err = updateCurrent(statesDirPath, jsonHash)
		if nil != err {
			return err
		}

		color.Green("Done.\n")
	} else {
		historyFile, err := os.Open(historyFilePath)
		if nil != err {
			return err
		}
		defer historyFile.Close()

		if string(jsonHash[:]) != handlers.GetLastState(historyFile) {
			color.Yellow("No database changes detected!\n")
		} else {
			historyFile.Close()

			historyFile, err := os.OpenFile(historyFilePath, os.O_WRONLY, 0644)
			if nil != err {
				return err
			}
			defer historyFile.Close()

			historyFile.WriteString(fmt.Sprintf("%x\n", jsonHash))

			err = updateCurrent(statesDirPath, jsonHash)
			if nil != err {
				return err
			}

			color.Green("Done.\n")
		}
	}

	return nil
}

func loadConfiguration() {
	fmt.Print("Loading configuration... ")

	config.LoadConfig()

	color.Green("Done.\n")
}

func updateCurrent(statesDirPath string, jsonHash [20]byte) error {
	currentStatePath := fmt.Sprintf("%x", jsonHash)
	currentFilePath := fmt.Sprintf("%s%c%s", statesDirPath, os.PathSeparator, "current")

	if _, err := os.Stat(currentFilePath); !os.IsNotExist(err) {
		os.Remove(currentFilePath)
	}

	err := os.Symlink(currentStatePath, currentFilePath)
	if nil != err {
		return err
	}

	return nil
}
