package command

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"time"

	"github.com/fatih/color"

	"github.com/c88lopez/dbs/common"
	"github.com/c88lopez/dbs/config"
	"github.com/c88lopez/dbs/entity"
	"github.com/c88lopez/dbs/help"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnPool *sql.DB

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

			startConnectionPool()
			s, err := buildSchemaState()
			if nil != err {
				return err
			}

			generateJsonSchemaState(s)
			if nil != err {
				return err
			}

			closeConnectionPool()
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

func startConnectionPool() {
	fmt.Print("Initializing connection pool... ")

	openConnectionPool(
		config.Parameters.Username + ":" + config.Parameters.Password + "@" + config.Parameters.Protocol + "(" +
			config.Parameters.Host + ":" + config.Parameters.Port + ")/" + config.Parameters.Database)

	color.Green("Done.\n")
}

func openConnectionPool(dsn string) error {
	var err error

	DbConnPool, err = sql.Open(config.Parameters.Driver, dsn)
	if nil != err {
		return err
	}

	return nil
}

func closeConnectionPool() error {
	return DbConnPool.Close()
}

func buildSchemaState() (*entity.Schema, error) {
	fmt.Print("Building schema state... ")
	schema := new(entity.Schema)

	schema.Name = config.Parameters.Database
	err := schema.LoadInformationSchema(DbConnPool)
	if nil != err {
		return nil, err
	}
	schema.FetchTables(DbConnPool)

	color.Green("Done.\n")

	return schema, nil
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
	jsonFilePath := fmt.Sprintf("%v/%x", statesDirPath, jsonHash)

	historyFilePath := fmt.Sprintf("%v/%v", statesDirPath, "history")
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

		color.Green("Done.\n")
	} else {
		historyFile, err := os.Open(historyFilePath)
		if nil != err {
			return err
		}
		defer historyFile.Close()

		last := false

		scanner := bufio.NewScanner(historyFile)
		for scanner.Scan() {
			if bytes.Contains(scanner.Bytes(), []byte(fmt.Sprintf("%s", jsonHash))) {
				last = true
			} else {
				last = false
			}
		}

		if last {
			color.Yellow("No database changes detected!\n")
		} else {
			historyFile.Close()

			historyFile, err := os.OpenFile(historyFilePath, os.O_WRONLY, 0644)
			if nil != err {
				return err
			}
			defer historyFile.Close()

			historyFile.WriteString(fmt.Sprintf("%x", jsonHash))

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
	currentStatePath := fmt.Sprintf("%v/%v", statesDirPath, jsonHash)
	currentFilePath := fmt.Sprintf("%v/%v", statesDirPath, "current")

	fmt.Println(currentStatePath)
	fmt.Println(currentFilePath)

	err := os.Symlink(currentStatePath, currentFilePath)
	if nil != err {
		return err
	}

	return nil
}
