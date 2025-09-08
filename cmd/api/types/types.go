package types

import "time"

type Quote struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID        int64     // unique value for each comment
	Content   string    // the comment data
	Author    string    // the person who wrote the comment
	CreatedAt time.Time // database timestamp
	Version   int32     // incremented on each update
}
