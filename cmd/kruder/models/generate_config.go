package models

// GenerateConfig is the incoming request to generate a config
type GenerateConfig struct {
	Language  string `json:"language"`
	Database  string `json:"database"`
	SchemaIDs []int  `json:"schemaIds"`
}
