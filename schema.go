package main

// Schema struct
type Schema struct {
	name          string
	charsetName   string
	collationName string

	tableCount int
	tables     []Table
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
	totalTables := 0

	return totalTables
}
