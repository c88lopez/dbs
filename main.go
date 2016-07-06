package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"crypto/sha1"

	_ "github.com/go-sql-driver/mysql"
)

var config = new(Config)
var dbConnPool *sql.DB

func main() {
	fmt.Println("** Welcome ton DBS **")
	fmt.Println("")

	loadConfiguration()

	startConnectionPool()
	s := buildSchemaState()
	generateJsonSchemaState(s)

	dbConnPool.Close()
}

func loadConfiguration() {
	fmt.Print("Loading configuration... ")

	config.loadConfig()

	fmt.Println("Done.")
}

func startConnectionPool() {
	fmt.Print("Initializing connection pool... ")

	var dsn string
	var err error

	dsn = config.getUsername() + ":" + config.getPassword() + "@/" + config.getDatabase()

	dbConnPool, err = sql.Open(config.getDriver(), dsn)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Done.")
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

	fmt.Printf("%s\n", string(schemaJson))

	hasher := sha1.New()
	hasher.Write([]byte(string(schemaJson)))
	bs := hasher.Sum(nil)

	fmt.Printf("%x\n", bs)

	fmt.Println("Done.")
}
