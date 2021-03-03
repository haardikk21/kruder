package models

// SchemaRequest is the incoming request to create a schema
type SchemaRequest struct {
	Name   string            `json:"name"`
	Fields map[string]string `json:"fields"`
}
