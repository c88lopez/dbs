package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := new(Config)
	config.loadConfig()

	db, err := sql.Open(config.getDriver(),
		config.getUsername()+":"+config.getPassword()+"@/"+config.getDatabase())

	if err != nil {
		log.Panic(err)
	}

	schema := new(Schema)

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	schema.SetConn(db)
	schema.FetchTables()

	fmt.Println(schema.GetTables())
}
