package entity

import "database/sql"

// Index struct
type Index struct {
	NonUnique    string         `json:"nonUnique"`
	KeyName      string         `json:"keyName"`
	SeqInIndex   string         `json:"seqInIndex"`
	ColumnName   string         `json:"columnName"`
	Collation    string         `json:"collation"`
	Cardinality  string         `json:"cardinality"`
	SubPart      sql.NullString `json:"subPart"`
	Packed       sql.NullString `json:"packed"`
	Null         string         `json:"null"`
	IndexType    string         `json:"indexType"`
	Comment      string         `json:"comment"`
	IndexComment string         `json:"indexComment"`
}
