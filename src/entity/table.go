package entity

import (
	"database/sql"

	"github.com/c88lopez/dbs/src/database"
)

// Table struct
type Table struct {
	Name string `json:"name"`

	Columns []Column `json:"columns"`
	Indexes []Index  `json:"indexes"`

	Engine         string `json:"engine"`
	DefaultCharset string `json:"defaultCharset"`
}

// AddColumn func
func (t *Table) AddColumn(c Column) *Table {
	t.Columns = append(t.Columns, c)

	return t
}

// AddIndex func
func (t *Table) AddIndex(i Index) *Table {
	t.Indexes = append(t.Indexes, i)

	return t
}

// FetchColumns func
func (t *Table) Fetch() error {
	var err error

	err = t.fetchColumns()
	if nil != err {
		return err
	}

	err = t.fetchIndexes()
	if nil != err {
		return err
	}

	return nil
}

// fetchColumns func
func (t *Table) fetchColumns() error {
	var result *sql.Rows

	result, err := DbConnPool.Query("DESCRIBE " + t.Name)
	if nil != err {
		return err
	}

	for result.Next() {
		var column Column

		result.Scan(&column.Name,
			&column.DataType,
			&column.Nullable,
			&column.Key,
			&column.DefaultValue,
			&column.Extra)

		t.AddColumn(column)
	}

	return nil
}

// fetchIndexes func
func (t *Table) fetchIndexes() error {
	var result *sql.Rows

	result, err := database.DbConnPool.Query("SHOW INDEX FROM " + t.Name)
	if nil != err {
		return err
	}

	for result.Next() {
		var index Index

		result.Scan(&index.NonUnique,
			&index.KeyName,
			&index.SeqInIndex,
			&index.ColumnName,
			&index.Collation,
			&index.Cardinality,
			&index.SubPart,
			&index.Packed,
			&index.Null,
			&index.IndexType,
			&index.Comment,
			&index.IndexComment)

		t.AddIndex(index)
	}

	return nil
}
