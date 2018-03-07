package entity

import "database/sql"

// Table struct
type Table struct {
	Name string `json:"name"`

	Columns     []Column     `json:"columns"`
	Indexes     []Index      `json:"indexes"`
	ForeignKeys []ForeignKey `json:"foreignKeys"`

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

// AddForeignKey func
func (t *Table) AddForeignKey(f ForeignKey) *Table {
	t.ForeignKeys = append(t.ForeignKeys, f)

	return t
}

// Fetch func
func (t *Table) Fetch(db *sql.DB) error {
	err := t.FetchColumns(db)
	if nil == err {
		err = t.FetchIndexes(db)
		if nil == err {
			err = t.FetchForeignKeys(db)
		}
	}

	return err
}

// FetchColumns func
func (t *Table) FetchColumns(db *sql.DB) error {
	var err error
	var result *sql.Rows

	result, err = db.Query("DESCRIBE " + t.Name)
	if nil == err {
		for result.Next() {
			var column Column

			err = result.Scan(&column.Name,
				&column.DataType,
				&column.Nullable,
				&column.Key,
				&column.DefaultValue,
				&column.Extra)

			if nil != err {
				return err
			}

			t.AddColumn(column)
		}
	}

	return nil
}

// FetchIndexes func
func (t *Table) FetchIndexes(db *sql.DB) error {
	var err error
	var result *sql.Rows

	result, err = db.Query("SHOW INDEX FROM " + t.Name)
	if nil != err {
		return err
	}

	var tableName string

	for result.Next() {
		var index Index

		err = result.Scan(&tableName,
			&index.NonUnique,
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

		if nil != err {
			return err
		}

		t.AddIndex(index)
	}

	return nil
}

// FetchForeignKeys func
func (t *Table) FetchForeignKeys(db *sql.DB) error {
	var err error
	var result *sql.Rows

	result, err = db.Query(`SELECT
	table_name,
		column_name,
		constraint_name,
		referenced_table_name,
		referenced_column_name
	FROM information_schema.key_column_usage
	WHERE referenced_column_name IS NOT NULL AND table_name = '` + t.Name + `'`)

	if nil == err {
		for result.Next() {
			var foreignKey ForeignKey

			err = result.Scan(&foreignKey.TableName,
				&foreignKey.ColumnName,
				&foreignKey.ConstraintName,
				&foreignKey.ReferencedTableName,
				&foreignKey.ReferencedColumnName)

			if nil != err {
				return err
			}

			t.AddForeignKey(foreignKey)
		}
	}

	return err
}
