package database

import (
	"fmt"
	"qotd/cmd/api/types"
)

var InMemoryQuotes []types.Quote
var InMemoryComments []types.Comment

type DatabaseType int

const (
	InMemory DatabaseType = iota
	Postgres
)

type Database struct {
	connectionString string
	dbType           DatabaseType
}

func NewDatabase(dbType DatabaseType, connectionString *string) *Database {
	// @todo Database connection logic
	if connectionString == nil {
		connectionString = new(string)
		*connectionString = ""
	}
	return &Database{
		connectionString: *connectionString,
		dbType:           dbType,
	}
}

func (db *Database) Connect() error {
	// @todo Implement database connection logic
	switch db.dbType {
	case InMemory:
		// Connect to in-memory database
		return nil
	case Postgres:
		// Connect to Postgres database
		return nil
	}
	return fmt.Errorf("unsupported database type")
}

func (db *Database) Disconnect() error {
	// @todo Implement database disconnection logic
	switch db.dbType {
	case InMemory:
		// Disconnect from in-memory database
		// Flush the in-memory quotes
		InMemoryQuotes = nil
		return nil
	case Postgres:
		// Disconnect from Postgres database
		return nil
	}
	return fmt.Errorf("unsupported database type")
}
