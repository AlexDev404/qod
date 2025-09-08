package database

import (
	"fmt"
	"qotd/cmd/api/types"
)

func (db *Database) GetComments() ([]types.Comment, error) {
	// @todo Implement fetching comments from the database
	switch db.dbType {
	case InMemory:
		return InMemoryComments, nil
	case Postgres:
		// Fetch comments from Postgres database
		return nil, nil
	}
	return nil, fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) WriteComment(comment types.Comment) error {
	// @todo Implement writing comments to the database
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
		// Find the last comment's ID and assign the next ID
		lastID := int64(0)
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
		return nil
	}
	return fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) GetCommentByID(id int64) (*types.Comment, error) {
	// @todo Implement fetching a single comment by ID from the database
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
		return nil, nil
	}
	return nil, fmt.Errorf(DATABASE_UNSUPPORTED)
}

func (db *Database) ModifyComment(commentID int64, comment types.Comment) error {
	// @todo Implement modifying a comment in the database
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
