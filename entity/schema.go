package entity

import "database/sql"

// Schema struct
type Schema struct {
	Name          string `json:"name"`
	CharsetName   string `json:"charsetName"`
	CollationName string `json:"collationName"`

	TableCount int     `json:"tableCount"`
	Tables     []Table `json:"tables"`
}

// LoadTables func
func (s *Schema) LoadTables() bool {
	return true
}

// AddTable func
func (s *Schema) AddTable(t Table) *Schema {
	s.Tables = append(s.Tables, t)
	s.TableCount++

	return s
}

// GetTables func
func (s *Schema) GetTables() []Table {
	return s.Tables
}

// LoadInformationSchema func
func (s *Schema) LoadInformationSchema(pool *sql.DB) error {
	rows, err := pool.Query("SELECT DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = '" + s.Name + "'")
	if nil != err {
		return err
	}
	defer rows.Close()

	rows.Next()

	var charset, collation string
	rows.Scan(&charset, &collation)

	s.CharsetName = charset
	s.CollationName = collation

	return nil
}

// FetchTables func
func (s *Schema) FetchTables(pool *sql.DB) (int, error) {
	s.TableCount = 0

	rows, err := pool.Query("SHOW TABLES")
	if nil != err {
		return -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var table Table

		if nil != rows.Scan(&table.Name) {
			return -1, err
		}

		table.FetchColumns(pool)

		s.AddTable(table)
	}

	return s.TableCount, nil
}
