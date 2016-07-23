package main

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
func (s *Schema) LoadInformationSchema() *Schema {
	rows, err := dbConnPool.Query("SELECT DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = '" + s.Name + "'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	rows.Next()

	var charset, collation string
	rows.Scan(&charset, &collation)

	s.CharsetName = charset
	s.CollationName = collation

	return s
}

// FetchTables func
func (s *Schema) FetchTables() int {
	s.TableCount = 0

	rows, err := dbConnPool.Query("SHOW TABLES")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table Table

		err = rows.Scan(&table.Name)
		if err != nil {
			panic(err)
		}

		table.FetchColumns()

		s.AddTable(table)
	}

	return s.TableCount
}
