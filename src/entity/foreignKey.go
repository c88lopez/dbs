package entity

// Index struct
type ForeignKey struct {
	TableName            string `json:"tableName"`
	ColumnName           string `json:"columnName"`
	ConstraintName       string `json:"constraintName"`
	ReferencedTableName  string `json:"referencedTableName "`
	ReferencedColumnName string `json:"referencedColumnName"`
}
