package entity

import "database/sql"

// Schema struct
type Schema struct {
	Name          string `json:"name"`
	CharsetName   string `json:"charsetName"`
	CollationName string `json:"collationName"`

	Tables []Table `json:"tables"`
}

// LoadTables func
func (s *Schema) LoadTables() bool {
	return true
}

// AddTable func
func (s *Schema) AddTable(t Table) *Schema {
	s.Tables = append(s.Tables, t)

	return s
}

// GetTables func
func (s *Schema) GetTables() []Table {
	return s.Tables
}

// LoadInformationSchema func
func (s *Schema) LoadInformationSchema(db *sql.DB) error {
	rows, err := db.Query(`SELECT DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME 
	FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = '` + s.Name + `'`)

	if nil == err {
		defer rows.Close()

		rows.Next()

		var charset, collation string
		rows.Scan(&charset, &collation)

		s.CharsetName = charset
		s.CollationName = collation
	}

	return err
}

// FetchTables func
func (s *Schema) FetchTables(db *sql.DB) error {
	var err error

	rows, err := db.Query("SHOW TABLES")
	if nil == err {
		defer rows.Close()

		for rows.Next() {
			var table Table

			if nil != rows.Scan(&table.Name) {
				return err
			}

			err = table.Fetch(db)
			if nil != err {
				return err
			}

			s.AddTable(table)
		}
	}

	return err
}
