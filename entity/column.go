package entity

// Column struct
type Column struct {
	Name         string `json:"name"`
	Datatype     string `json:"datatype"`
	Nullable     string `json:"nullable"`
	IsKey        string `json:"isKey"`
	DefaultValue string `json:"defaultValue"`
	Extra        string `json:"extra"`
}
