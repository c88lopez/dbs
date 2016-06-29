package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var config = new(Config)
var dbConnPool *sql.DB

func main() {
	loadConfiguration()

	startConnectionPool()
	buildSchemaState()

	dbConnPool.Close()
}

func loadConfiguration() {
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
}

func buildSchemaState() {
	schema := new(Schema)

	schema.SetConn(dbConnPool)
	schema.SetName(config.getDatabase())
	schema.LoadInformationSchema().FetchTables()
}
