package entity

import "database/sql"

// Table struct
type Table struct {
	Name string `json:"name"`

	Columns []Column `json:"columns"`

	Engine         string `json:"engine"`
	DefaultCharset string `json:"defaultCharset"`
}

// AddColumn func
func (t *Table) AddColumn(c Column) *Table {
	t.Columns = append(t.Columns, c)

	return t
}

// FetchColumns func
func (t *Table) FetchColumns(pool *sql.DB) (*Table, error) {
	var result *sql.Rows

	result, err := pool.Query("DESCRIBE " + t.Name)
	if nil != err {
		return nil, err
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

	return t, nil
}
