package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"

	"time"

	"os"

	"bufio"

	"bytes"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

var dbConnPool *sql.DB

func main() {
	start := time.Now()

	if len(os.Args) == 1 {
		showHelp()
	} else {
		switch os.Args[1] {
		default:
			showHelp()
		case "init":
			generateInitFolder()
			break
		case "config":
			setDatabaseConfigInteractively()
			break
		case "generate":
			setConfigFilePath()
			loadConfiguration()

			startConnectionPool()
			s := buildSchemaState()
			generateJsonSchemaState(s)

			dbConnPool.Close()
			break
		}
	}

	fmt.Printf("\nElapsed time: %s\n", time.Since(start))
}

func loadConfiguration() {
	fmt.Print("Loading configuration... ")

	config.loadConfig()

	color.Green("Done.\n")
}

func startConnectionPool() {
	fmt.Print("Initializing connection pool... ")

	var dsn string
	var err error

	dsn = config.Username + ":" + config.Password + "@/" + config.Database

	dbConnPool, err = sql.Open(config.Driver, dsn)
	check(err)

	color.Green("Done.\n")
}

func buildSchemaState() *Schema {
	fmt.Print("Building schema state... ")
	schema := new(Schema)

	schema.Name = config.Database
	schema.LoadInformationSchema().FetchTables()

	color.Green("Done.\n")

	return schema
}

func generateJsonSchemaState(s *Schema) {
	fmt.Print("Generating json and json hash... ")

	var err error

	schemaJson, err := json.Marshal(s)
	check(err)

	jsonHash := sha1.Sum(schemaJson)

	statesDirPath := getStatesDirPath()
	jsonFilePath := fmt.Sprintf("%v/%x", statesDirPath, jsonHash)

	historyFilePath := fmt.Sprintf("%v/%v", statesDirPath, "history")
	if _, err = os.Stat(jsonFilePath); os.IsNotExist(err) {
		jsonFile, err := os.Create(jsonFilePath)
		check(err)
		defer jsonFile.Close()

		_, err = jsonFile.Write(schemaJson)
		check(err)

		historyFile, err := os.OpenFile(historyFilePath, os.O_WRONLY|os.O_APPEND, 0644)
		check(err)
		defer historyFile.Close()

		historyFile.WriteString(fmt.Sprintf("%x\n", jsonHash))

		color.Green("Done.\n")
	} else {
		historyFile, err := os.Open(historyFilePath)
		check(err)
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

			historyFile, err := os.OpenFile(historyFilePath, os.O_WRONLY|os.O_APPEND, 0644)
			check(err)
			defer historyFile.Close()

			historyFile.WriteString(fmt.Sprintf("%x\n", jsonHash))

			color.Green("Done.\n")
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
