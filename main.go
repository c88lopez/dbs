package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:qweqwe@/dbs")

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
