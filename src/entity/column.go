package entity

import "database/sql"

// Column struct
type Column struct {
	Name         string         `json:"name"`
	DataType     string         `json:"datatype"`
	Nullable     string         `json:"nullable"`
	Key          string         `json:"key"`
	DefaultValue sql.NullString `json:"defaultValue"`
	Extra        string         `json:"extra"`
}
