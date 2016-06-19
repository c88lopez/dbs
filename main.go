package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbConnPool *sql.DB

func main() {
	var err error

	config := new(Config)
	config.loadConfig()

	dbConnPool, err = sql.Open(config.getDriver(),
		config.getUsername()+":"+config.getPassword()+"@/"+config.getDatabase())

	if err != nil {
		log.Panic(err)
	}

	buildSchemaState(config)
}

func buildSchemaState(config *Config) {
	schema := new(Schema)

	schema.SetConn(dbConnPool)
	schema.SetName(config.getDatabase())
	schema.LoadInformationSchema().FetchTables()
}
