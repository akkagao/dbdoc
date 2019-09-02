package models

type TableField struct {
	Field      string `json:"field"`
	Type       string `json:"type"`
	Collation  string `json:"collation"`
	Null       string `json:"null"`
	Key        string `json:"key"`
	Default    string `json:"default"`
	Extra      string `json:"extra"`
	Privileges string `json:"privileges"`
	Comment    string `json:"comment"`
}
