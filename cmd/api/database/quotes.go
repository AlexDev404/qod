package database

import (
	"fmt"
	"qotd/cmd/api/types"
)

func (db *Database) GetQuotes() ([]types.Quote, error) {
	// @todo Implement fetching quotes from the database
	switch db.dbType {
	case InMemory:
		return InMemoryQuotes, nil
	case Postgres:
		// Fetch quotes from Postgres database
		return nil, nil
	}
	return nil, fmt.Errorf("unsupported database type")
}

func (db *Database) WriteQuote(quote types.Quote) error {
	// @todo Implement writing quotes to the database
	switch db.dbType {
	case InMemory:
		InMemoryQuotes = append(InMemoryQuotes, quote)
		return nil
	case Postgres:
		// Write quote to Postgres database
		return nil
	}
	return fmt.Errorf("unsupported database type")
}

func (db *Database) GetQuoteByID(id int) (*types.Quote, error) {
	// @todo Implement fetching a single quote by ID from the database
	switch db.dbType {
	case InMemory:
		for _, quote := range InMemoryQuotes {
			if quote.ID == id {
				return &quote, nil
			}
		}
		return nil, fmt.Errorf("quote not found")
	case Postgres:
		// Fetch quote by ID from Postgres database
		return nil, nil
	}
	return nil, fmt.Errorf("unsupported database type")
}

func (db *Database) ModifyQuote(quoteID int, quote types.Quote) error {
	// @todo Implement modifying a quote in the database
	switch db.dbType {
	case InMemory:
		for i, q := range InMemoryQuotes {
			if q.ID == quoteID {
				InMemoryQuotes[i] = quote
				return nil
			}
		}
		return fmt.Errorf("quote not found")
	case Postgres:
		// Modify quote in Postgres database
		return nil
	}
	return fmt.Errorf("unsupported database type")
}
