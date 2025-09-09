package types

import "time"

type Quote struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID        int       `json:"id"`         // unique value for each comment
	Content   string    `json:"content"`    // the comment data
	Author    string    `json:"author"`     // the person who wrote the comment
	CreatedAt time.Time `json:"created_at"` // database timestamp
	Version   int       `json:"version"`    // incremented on each update
}
