package main

import (
	"database/sql"
	"fmt"
	"log"
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

	schema.SetConn(db)
	schema.FetchTables()

	fmt.Println("vamo lo pibe")
}
