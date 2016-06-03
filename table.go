package main

// Table struct
type Table struct {
	name string

	columns []Column
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

	return t
}
