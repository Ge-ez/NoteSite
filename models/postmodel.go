package models

import "time"

type Post struct {
	ID      uint64
	Title   string
	Content string

	AuthorID  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
