package models

// Endpoint is a Kruder endpoint
type Endpoint struct {
	ID       int      `json:"id"`
	Path     string   `json:"path"`
	SchemaID int      `json:"schemaId"`
	Fields   []string `json:"fields"`
}
