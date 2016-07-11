package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"crypto/sha1"

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
		case "build-schema-state":
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
	fmt.Printf("Loading configuration... ")

	config.loadConfig()

	color.Green("Done.\n")
}

func startConnectionPool() {
	fmt.Printf("Initializing connection pool... ")

	var dsn string
	var err error

	dsn = config.getUsername() + ":" + config.getPassword() + "@/" + config.getDatabase()

	dbConnPool, err = sql.Open(config.getDriver(), dsn)
	if err != nil {
		log.Panic(err)
	}

	color.Green("Done.\n")
}

func buildSchemaState() *Schema {
	fmt.Print("Building schema state... ")
	schema := new(Schema)

	schema.SetName(config.getDatabase())
	schema.LoadInformationSchema().FetchTables()

	color.Green("Done.\n")

	return schema
}

func generateJsonSchemaState(s *Schema) {
	fmt.Print("Generating json and json hash... ")

	schemaJson, err := json.Marshal(s)
	if err != nil {
		log.Panic(err)
	}

	jsonHash := generateJsonSchemaStateHash(schemaJson)

	statesDirPath := getStatesDirPath()
	jsonFilePath := statesDirPath + "/" + jsonHash
	jsonFile, err := os.Create(jsonFilePath)
	if err != nil {
		log.Panic(err)
	}

	_, err = jsonFile.Write(schemaJson)
	if err != nil {
		log.Panic(err)
	}

	color.Green("Done.\n")
}

func generateJsonSchemaStateHash(schemaJson []byte) string {
	hasher := sha1.New()
	_, err := hasher.Write(schemaJson)
	if err != nil {
		log.Panic(err)
	}

	return string(hasher.Sum(nil))
}
