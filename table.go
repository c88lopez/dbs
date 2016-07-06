package main

import (
	"database/sql"
	"log"
)

// Table struct
type Table struct {
	Name string `json:"name"`

	Columns []Column `json:"columns"`

	Engine         string `json:"engine"`
	DefaultCharset string `json:"defaultCharset"`
}

// GetConn func
func (t *Table) GetConn() *sql.DB {
	return dbConnPool
}

// SetName func
func (t *Table) SetName(name string) *Table {
	t.Name = name

	return t
}

// GetName func
func (t *Table) GetName() string {
	return t.Name
}

// AddColumn func
func (t *Table) AddColumn(c Column) *Table {
	t.Columns = append(t.Columns, c)

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

		result.Scan(&column.Name,
			&column.Datatype,
			&column.Nullable,
			&column.IsKey,
			&column.DefaultValue,
			&column.Extra)

		t.AddColumn(column)
	}

	return t
}
