package database

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"

	"github.com/c88lopez/dbs/src/entity"
	"github.com/c88lopez/dbs/src/handlers"
	"github.com/c88lopez/dbs/src/mainfolder"

	"github.com/fatih/color"
)

// New new
func New() error {
	if err := StartConnectionPool(); nil != err {
		return err
	}
	defer CloseConnectionPool()

	s, err := BuildSchemaState()
	if nil != err {
		return err
	}

	return generateJSONSchemaState(s)
}

// generateJSONSchemaState(s *entity.Schema) generateJSONSchemaState
func generateJSONSchemaState(s *entity.Schema) error {
	fmt.Print("Generating json and json hash... ")

	schemaJSON, err := json.Marshal(s)
	if nil != err {
		return err
	}

	jsonHash := fmt.Sprintf("%x", sha1.Sum(schemaJSON))

	statesDirPath := mainfolder.GetStatesDirPath()
	jsonFilePath := fmt.Sprintf("%v%c%s", statesDirPath, os.PathSeparator, jsonHash)

	historyFilePath := fmt.Sprintf("%v%c%v", statesDirPath, os.PathSeparator, "history")
	if _, err = os.Stat(jsonFilePath); os.IsNotExist(err) {
		if err = createJSONFile(jsonFilePath, schemaJSON); err != nil {
			return err
		}

		if err = updateHistoryFile(mainfolder.GetHistoryFilePath(), jsonHash); err != nil {
			return err
		}

		if err = updateCurrent(statesDirPath, jsonHash); nil != err {
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

			if err = updateHistoryFile(historyFilePath, jsonHash); err != nil {
				return err
			}

			if err = updateCurrent(statesDirPath, jsonHash); nil != err {
				return err
			}

			color.Green("Done.\n")
		}
	}

	return nil
}

func updateCurrent(statesDirPath, jsonHash string) error {
	currentFilePath := fmt.Sprintf("%s%c%s", statesDirPath, os.PathSeparator, "current")

	if _, err := os.Stat(currentFilePath); !os.IsNotExist(err) {
		os.Remove(currentFilePath)
	}

	return os.Symlink(jsonHash, currentFilePath)
}

func createJSONFile(jsonFilePath string, schemaJSON []byte) error {
	jsonFile, err := os.Create(jsonFilePath)
	if nil != err {
		return err
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(schemaJSON)

	return err
}

func updateHistoryFile(historyFilePath, jsonHash string) error {
	historyFile, err := os.OpenFile(historyFilePath, os.O_WRONLY|os.O_APPEND, 0644)
	if nil != err {
		return err
	}
	defer historyFile.Close()

	_, err = historyFile.WriteString(jsonHash)

	return err
}
