package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var config = new(Config)
var dbConnPool *sql.DB

func main() {
	fmt.Println("** Welcome ton DBS **")
	fmt.Println("")

	loadConfiguration()

	startConnectionPool()
	buildSchemaState()

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

func buildSchemaState() {
	schema := new(Schema)

	schema.SetConn(dbConnPool)
	schema.SetName(config.getDatabase())
	schema.LoadInformationSchema().FetchTables()
}
