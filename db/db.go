package db

import (
	_ "embed"
)

//go:embed schema.sql
var schema []byte

func GetSchema() string {
	return string(schema)
}
