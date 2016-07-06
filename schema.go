package main

import (
	"database/sql"
	"log"
)

// Schema struct
type Schema struct {
	Name          string `json:"name"`
	CharsetName   string `json:"charsetName"`
	CollationName string `json:"collationName"`

	TableCount int     `json:"tableCount"`
	Tables     []Table `json:"tables"`
}

// GetConn func
func (s *Schema) GetConn() *sql.DB {
	return dbConnPool
}

// SetName func
func (s *Schema) SetName(name string) *Schema {
	s.Name = name

	return s
}

// GetName func
func (s *Schema) GetName() string {
	return s.Name
}

// SetCharsetName func
func (s *Schema) SetCharsetName(charsetName string) *Schema {
	s.CharsetName = charsetName

	return s
}

// GetCharsetName func
func (s *Schema) GetCharsetName() string {
	return s.CharsetName
}

// SetCollationName func
func (s *Schema) SetCollationName(collationName string) *Schema {
	s.CollationName = collationName

	return s
}

// GetCollationName func
func (s *Schema) GetCollationName() string {
	return s.CollationName
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
	rows, err := s.GetConn().Query("SELECT DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = '" + s.GetName() + "'")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	rows.Next()

	var charset, collation string
	rows.Scan(&charset, &collation)

	s.SetCharsetName(charset).SetCollationName(collation)

	return s
}

// FetchTables func
func (s *Schema) FetchTables() int {
	s.TableCount = 0

	rows, err := s.GetConn().Query("SHOW TABLES")

	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table Table

		err = rows.Scan(&table.Name)
		if err != nil {
			log.Panic(err)
		}

		table.FetchColumns()

		s.AddTable(table)
	}

	return s.TableCount
}
