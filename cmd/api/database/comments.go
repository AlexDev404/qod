package database

import (
	"context"
	"fmt"
	"qotd/cmd/api/types"
	"time"
)

func (db *Database) GetComments() ([]types.Comment, error) {
	// @completed Implement fetching comments from the database
	switch db.dbType {
	case InMemory:
		return InMemoryComments, nil
	case Postgres:
		// Fetch comments from Postgres database
		query := `
			SELECT id, content, author, created_at, version
			FROM comments
		`
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		rows, err := db.context.QueryContext(ctx, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var comments []types.Comment
		for rows.Next() {
			var c types.Comment
			if err := rows.Scan(&c.ID, &c.Content, &c.Author, &c.CreatedAt, &c.Version); err != nil {
				return nil, err
			}
			comments = append(comments, c)
		}
		return comments, nil
	}
	return nil, fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) WriteComment(comment types.Comment) error {
	// @completed Implement writing comments to the database
	/*
	 * Already completed
	 * ===============
	 * 1. In-memory database writing
	 * 2. PostgreSQL writing
	 */
	switch db.dbType {
	case InMemory:
		// Find the last comment's ID and assign the next ID
		lastID := int(0)
		for _, c := range InMemoryComments {
			if c.ID > lastID {
				lastID = c.ID
			}
		}
		comment.ID = lastID + 1
		InMemoryComments = append(InMemoryComments, comment)
		return nil
	case Postgres:
		// Write comment to Postgres database
		query := `
			INSERT INTO comments (content, author)
			VALUES ($1, $2)
			RETURNING id, created_at, version
		`
		args := []any{comment.Content, comment.Author}
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		err := db.context.QueryRowContext(ctx, query, args...).Scan(&comment.ID, &comment.CreatedAt, &comment.Version)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) GetCommentByID(id int) (*types.Comment, error) {
	// @completed Implement fetching a single comment by ID from the database
	switch db.dbType {
	case InMemory:
		for _, comment := range InMemoryComments {
			if comment.ID == id {
				return &comment, nil
			}
		}
		return nil, fmt.Errorf("comment not found")
	case Postgres:
		// Fetch comment by ID from Postgres database
		query := `
			SELECT id, content, author, created_at, version
			FROM comments
			WHERE id = $1
		`
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		var c types.Comment
		err := db.context.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Content, &c.Author, &c.CreatedAt, &c.Version)
		if err != nil {
			return nil, err
		}
		return &c, nil
	}
	return nil, fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) ModifyComment(commentID int, comment types.Comment) error {
	// @completed Implement modifying a comment in the database
	switch db.dbType {
	case InMemory:
		for i, c := range InMemoryComments {
			if c.ID == commentID {
				InMemoryComments[i] = comment
				return nil
			}
		}
		return fmt.Errorf("comment not found")
	case Postgres:
		// Modify comment in Postgres database
		query := `
			UPDATE comments
			SET content = $1, author = $2, created_at = $3, version = version + 1
			WHERE id = $4
			RETURNING version
		`
		args := []any{comment.Content, comment.Author, time.Now(), commentID}
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		err := db.context.QueryRowContext(ctx, query, args...).Scan(&comment.Version)
		if err != nil {
			return err
		}

		return nil
	}
	return fmt.Errorf(DATABASE_UNSUPPORTED)
}

func ValidateComment(comment types.Comment) error {
	if comment.Author == "" {
		return fmt.Errorf("Field 'Author' missing")
	}

	if comment.Content == "" {
		return fmt.Errorf("Field 'Content' missing")
	}
	return nil
}

func (db *Database) DeleteComment(id int) error {
	// @completed Implement deleting a comment from the database
	switch db.dbType {
	case InMemory:
		for i, c := range InMemoryComments {
			if c.ID == id {
				InMemoryComments = append(InMemoryComments[:i], InMemoryComments[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("comment not found")
	case Postgres:
		// Delete comment from Postgres database
		query := `
			DELETE FROM comments
			WHERE id = $1
		`
		ctx, cancel := context.WithTimeout(context.Background(), db.queryTimeout)
		defer cancel()

		result, err := db.context.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return fmt.Errorf("comment not found")
		}
		return nil
	}
	return fmt.Errorf(DATABASE_UNSUPPORTED)
}
