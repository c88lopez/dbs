package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"crypto/sha1"

	_ "github.com/go-sql-driver/mysql"
	"time"
	"os"
)

var config = new(Config)
var dbConnPool *sql.DB

func main() {
	start := time.Now()

	fmt.Printf("** Welcome to DBS **\n")
	fmt.Printf("Version 0.0.1\n\n")

	fmt.Printf("%#v", os.Args[1:])

	loadConfiguration()

	startConnectionPool()
	s := buildSchemaState()
	generateJsonSchemaState(s)

	dbConnPool.Close()

	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}

func loadConfiguration() {
	fmt.Printf("Loading configuration... ")

	config.loadConfig()

	fmt.Printf("Done.\n")
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

	fmt.Print("Done.\n")
}

func buildSchemaState() *Schema {
	fmt.Print("Building schema state... ")
	schema := new(Schema)

	schema.SetName(config.getDatabase())
	schema.LoadInformationSchema().FetchTables()

	fmt.Println("Done.")

	return schema
}

func generateJsonSchemaState(s *Schema) {
	fmt.Print("Generating json... ")

	schemaJson, err := json.Marshal(s)
	if err != nil {
		log.Panic(err)
	}

	//fmt.Printf("%s\n", string(schemaJson))

	hasher := sha1.New()
	hasher.Write([]byte(string(schemaJson)))
	bs := hasher.Sum(nil)

	fmt.Printf("%x\n", bs)

	fmt.Printf("Done.\n")
}
