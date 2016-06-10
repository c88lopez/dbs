package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Schema struct
type Schema struct {
	conn *sql.DB

	name          string
	charsetName   string
	collationName string

	tableCount int
	tables     []Table
}

// SetConn func
func (s *Schema) SetConn(conn *sql.DB) *Schema {
	s.conn = conn

	return s
}

// GetConn func
func (s *Schema) GetConn() *sql.DB {
	return s.conn
}

// SetName func
func (s *Schema) SetName(name string) *Schema {
	s.name = name

	return s
}

// GetName func
func (s *Schema) GetName() string {
	return s.name
}

// SetCharsetName func
func (s *Schema) SetCharsetName(charsetName string) *Schema {
	s.charsetName = charsetName

	return s
}

// GetCharsetName func
func (s *Schema) GetCharsetName() string {
	return s.charsetName
}

// SetCollationName func
func (s *Schema) SetCollationName(collationName string) *Schema {
	s.collationName = collationName

	return s
}

// GetCollationName func
func (s *Schema) GetCollationName() string {
	return s.collationName
}

// LoadTables func
func (s *Schema) LoadTables() bool {
	return true
}

// AddTable func
func (s *Schema) AddTable(t *Table) *Schema {
	return s
}

// GetTables func
func (s *Schema) GetTables() []Table {
	return s.tables
}

// FetchTables func
func (s *Schema) FetchTables() int {
	s.tableCount = 0

	query := "SHOW TABLES;"
	fmt.Println(query)

	rows, err := s.GetConn().Query(query)

	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var table Table

		err = rows.Scan(&table.name)
		if err != nil {
			log.Panic(err)
		}

		query = "SHOW CREATE TABLE " + table.GetName() + ";"
		fmt.Println(query)

		var result *sql.Rows
		result, err = s.GetConn().Query(query)
		if err != nil {
			log.Panic(err)
		}

		type TableRaw struct {
			a string
			b string
		}

		var tr TableRaw

		for result.Next() {
			result.Scan(&tr.a, &tr.b)
			fmt.Printf("%#v\n", tr)
		}

		s.tables = append(s.tables, table)

		s.tableCount++
	}

	return s.tableCount
}
