package models
import "time"
type Comment struct {
	ID         uint
	PostID     uint
	ReplierID  uint
	Reply     string
	CreatedAt time.Time
}