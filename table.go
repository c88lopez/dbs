package main

import (
	"database/sql"
	"log"
)

// Table struct
type Table struct {
	conn *sql.DB

	name string

	columns []Column

	engine         string
	defaultCharset string
}

// SetConn func
func (t *Table) SetConn(conn *sql.DB) *Table {
	t.conn = conn

	return t
}

// GetConn func
func (t *Table) GetConn() *sql.DB {
	return t.conn
}

// SetName func
func (t *Table) SetName(name string) *Table {
	t.name = name

	return t
}

// GetName func
func (t *Table) GetName() string {
	return t.name
}

// AddColumn func
func (t *Table) AddColumn(c Column) *Table {
	t.columns = append(t.columns, c)

	return t
}

// FetchColumns func
func (t *Table) FetchColumns() *Table {
	query := "DESCRIBE " + t.GetName()

	var result *sql.Rows
	result, err := t.GetConn().Query(query)
	if err != nil {
		log.Panic(err)
	}

	for result.Next() {
		var column Column

		result.Scan(&column.name,
			&column.datatype,
			&column.nullable,
			&column.isKey,
			&column.defaultValue,
			&column.extra)

		t.AddColumn(column)
	}

	return t
}
