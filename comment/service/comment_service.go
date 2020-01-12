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
//Updatecomment updates a given comment
func (as *CommentService) UpdateComment(comment *models.Comment) (*models.Comment, []error) {
	commt, errs := as.commentRepo.UpdateComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return commt, errs
}

//Deletecomment deletes a given comment
func (as *CommentService) DeleteComment(id uint) (*models.Comment, []error) {
	commt, errs := as.commentRepo.DeleteComment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return commt, errs
}

//StoreComment stores a given Comment
func (as *CommentService) StoreComment(comment *models.Comment) (*models.Comment, []error) {
	commt, errs := as.commentRepo.StoreComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return commt, errs
}

//Postcomments returns comments for a Post
func (as *CommentService) PostComments(post *models.Post, comment *models.Comment) ([]models.Comment, []error) {
	commts, errs := as.commentRepo.PostComments(post, comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return commts, errs
}
