package main

// Column struct
type Column struct {
	name string

	dataType string

	nullable      bool
	defaultValue  string
	autoincrement bool

	comment string
}
