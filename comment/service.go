package comment

import (
	"github.com/notesite/models"
)

//CommentService represents user - Comment operations
type CommentService interface {
	Comments() ([]models.Comment, []error)
	Comment(id uint) (*models.Comment, []error)
	PostComments(post *models.Post, comment *models.Comment) ([]models.Comment, []error)
	UpdateComment(comment *models.Comment) (*models.Comment, []error)
	DeleteComment(id uint) (*models.Comment, []error)
	StoreComment(comment *models.Comment) (*models.Comment, []error)
}
