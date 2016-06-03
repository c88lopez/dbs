package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:qweqwe@/db_algo")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(db)

	err = db.Ping()

	if err != nil {
		log.Panic(err)
	}

	rows, err := db.Query("SHOW TABLES")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(rows)

	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		if err != nil {
			log.Panic(err)
		}

		fmt.Println(name)
	}

	schema := new(Schema)
	fmt.Println(schema.GetTables())

	schema.SetName("trolo")
	schema.SetCharsetName("utf-8")
	schema.SetCollationName("unCollation")

	fmt.Println(schema)
}
