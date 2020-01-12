package comment

import (
	"github.com/notesite/models"
)

//CommentRepository represents user - Comment operations
type CommentRepository interface {
	Comments() ([]models.Comment, []error)
	Comment(id uint) (*models.Comment, []error)
	PostComments(post *models.Post, Comment *models.Comment) ([]models.Comment, []error)
	UpdateComment(Comment *models.Comment) (*models.Comment, []error)
	DeleteComment(id uint) (*models.Comment, []error)
	StoreComment(Comment *models.Comment) (*models.Comment, []error)
}
