package service

import (
	"github.com/notesite/comment"
	"github.com/notesite/models"
)

//commentService implements comment.Service
type CommentService struct {
	commentRepo comment.CommentRepository
}

// NewCommentService returns a new CommentService object
func NewCommentService(ansRepo comment.CommentRepository) comment.CommentService {
	return &CommentService{commentRepo: ansRepo}
}

//comments returns all stored comments
func (as *CommentService) Comments() ([]models.Comment, []error) {
	commts, errs := as.commentRepo.Comments()
	if len(errs) > 0 {
		return nil, errs
	}
	return commts, errs
}

//Comment retrieves stored Comment given its id
func (as *CommentService) Comment(id uint) (*models.Comment, []error) {
	commt, errs := as.commentRepo.Comment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return commt, errs
}
