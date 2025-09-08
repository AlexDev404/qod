package database

import (
	"context"
	"fmt"
	"qotd/cmd/api/types"
)

func (db *Database) GetQuotes() ([]types.Quote, error) {
	// @completed Implement fetching quotes from the database
	switch db.dbType {
	case InMemory:
		return InMemoryQuotes, nil
	case Postgres:
		// Fetch quotes from Postgres database
		query := `
			SELECT id, text, author, created_at
			FROM quotes
		`
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		rows, err := db.context.QueryContext(ctx, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var quotes []types.Quote
		for rows.Next() {
			var q types.Quote
			if err := rows.Scan(&q.ID, &q.Text, &q.Author, &q.CreatedAt); err != nil {
				return nil, err
			}
			quotes = append(quotes, q)
		}
		return quotes, nil
	}
	return nil, fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) WriteQuote(quote types.Quote) error {
	// @completed Implement writing quotes to the database
	/*
	 * Already completed
	 * ===============
	 * 1. In-memory database writing
	 *
	 * Not_completed_Todos
	 * ==============
	 * 1. PostgreSQL writing
	 */
	switch db.dbType {
	case InMemory:
		// Find the last quote's ID and assign the next ID
		lastID := 0
		for _, q := range InMemoryQuotes {
			if q.ID > lastID {
				lastID = q.ID
			}
		}
		quote.ID = lastID + 1
		InMemoryQuotes = append(InMemoryQuotes, quote)
		return nil
	case Postgres:
		// Write quote to Postgres database
		query := `
			INSERT INTO quotes (text, author)
			VALUES ($1, $2)
			RETURNING id, created_at
		`
		args := []any{quote.Text, quote.Author}
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		err := db.context.QueryRowContext(ctx, query, args...).Scan(&quote.ID, &quote.CreatedAt)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) GetQuoteByID(id int) (*types.Quote, error) {
	// @completed Implement fetching a single quote by ID from the database
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
		query := `
			SELECT id, text, author, created_at
			FROM quotes
			WHERE id = $1
		`
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		var q types.Quote
		err := db.context.QueryRowContext(ctx, query, id).Scan(&q.ID, &q.Text, &q.Author, &q.CreatedAt)
		if err != nil {
			return nil, err
		}
		return &q, nil
	}
	return nil, fmt.Errorf(DATABASE_UNSUPPORTED)
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
		query := `
			UPDATE quotes
			SET text = $1, author = $2
			WHERE id = $3
		`
		args := []any{quote.Text, quote.Author, quoteID}
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()
		_, err := db.context.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf(DATABASE_UNSUPPORTED)
}

func ValidateQuote(quote types.Quote) error {
	if quote.Author == "" {
		return fmt.Errorf("Field 'Author' missing")
	}

	if quote.Text == "" {
		return fmt.Errorf("Field 'Text' missing")
	}
	return nil
}
