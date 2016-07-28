package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"

	"time"

	"os"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

var dbConnPool *sql.DB

func main() {
	start := time.Now()

	displayHeader()

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

	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}

func displayHeader() {

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
	if err != nil {
		panic(err)
	}

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
	if err != nil {
		panic(err)
	}

	jsonHash := sha1.Sum(schemaJson)

	fmt.Sprintf("%x", jsonHash)

	statesDirPath := getStatesDirPath()
	jsonFilePath := fmt.Sprintf("%v/%x", statesDirPath, jsonHash)

	if _, err = os.Stat(jsonFilePath); os.IsNotExist(err) {
		jsonFile, err := os.Create(jsonFilePath)
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()

		_, err = jsonFile.Write(schemaJson)
		if err != nil {
			panic(err)
		}

		historyFilePath := fmt.Sprintf("%v/%v", statesDirPath, "history")
		historyFile, err := os.OpenFile(historyFilePath, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer historyFile.Close()

		historyFile.WriteString(fmt.Sprintf("%x\n", jsonHash))

		color.Green("Done.\n")
	} else {
		color.Yellow("No database changes detected!\n")
	}
}
