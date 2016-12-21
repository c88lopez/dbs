package command

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"

	"time"

	"github.com/fatih/color"

	"github.com/c88lopez/dbs/src/common"
	"github.com/c88lopez/dbs/src/config"
	"github.com/c88lopez/dbs/src/entity"
	"github.com/c88lopez/dbs/src/help"

	"github.com/c88lopez/dbs/src/database"
	"github.com/c88lopez/dbs/src/handlers"
	_ "github.com/go-sql-driver/mysql"
)

// Executes the dbs.
func Execute() error {
	start := time.Now()

	if len(os.Args) == 1 {
		help.ShowHelp()
	} else {
		switch os.Args[1] {
		default:
			help.ShowHelp()
		case "init":
			common.GenerateMainFolder()
			break
		case "config":
			config.SetDatabaseConfigInteractively()
			break
		case "generate":
			config.SetConfigFilePath()
			loadConfiguration()

			database.StartConnectionPool()
			defer database.CloseConnectionPool()

			s, err := database.BuildSchemaState()
			if nil != err {
				return err
			}

			generateJsonSchemaState(s)
			if nil != err {
				return err
			}

			break
		}
	}

	fmt.Printf("\nElapsed time: %s\n", time.Since(start))
	return nil
}

func loadConfiguration() {
	fmt.Print("Loading configuration... ")

	config.LoadConfig()

	color.Green("Done.\n")
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
