package entity

// Column struct
type Column struct {
	Name         string `json:"name"`
	DataType     string `json:"datatype"`
	Nullable     string `json:"nullable"`
	Key          string `json:"isKey"`
	DefaultValue string `json:"defaultValue"`
	Extra        string `json:"extra"`
}
