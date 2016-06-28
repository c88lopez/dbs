package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var config Config
var dbConnPool *sql.DB

func main() {
	loadConfiguration()
	startConnectionPool()
	buildSchemaState()
}

func loadConfiguration() {
	config := new(Config)
	config.loadConfig()
}

func startConnectionPool() {
	var dsn string
	var err error

	dsn = config.getUsername()
	dsn += ":" + config.getPassword()
	dsn += "@/" + config.getDatabase()

	dbConnPool, err = sql.Open(config.getDriver(), dsn)

	if err != nil {
		log.Panic(err)
	}

	schema := new(Schema)

	schema.SetConn(db)
	schema.FetchTables()

	fmt.Println("vamo lo pibe")
}

func buildSchemaState() {
	schema := new(Schema)

	schema.SetConn(dbConnPool)
	schema.SetName(config.getDatabase())
	schema.LoadInformationSchema().FetchTables()
}
