package models

import "time"

type Post struct {
	ID         uint
    UserID     uint
   	Text       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
func (Post) TableName() string {
	return "posts"
}