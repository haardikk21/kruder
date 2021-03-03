package models

// FieldDataType is all the data types that Kruder supports
type FieldDataType int

const (
	// String type
	String FieldDataType = iota
	// Number type
	Number
	// Boolean type
	Boolean
	// Array type
	Array
	// Object type
	Object
)

// Schema is a Kruder schema
type Schema struct {
	ID     int                      `json:"id"`
	Name   string                   `json:"name"`
	Fields map[string]FieldDataType `json:"fields"`
}
