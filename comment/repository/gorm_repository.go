package repository

import (
	"github.com/notesite/comment"
	"github.com/notesite/models"
	"github.com/jinzhu/gorm"
)

//CommentGormRepo implements comment.CommentRepository interface
type CommentGormRepo struct {
	conn *gorm.DB
}

//NewCommentGormRepo returns new object of CommentGormRepo
func NewCommentGormRepo(db *gorm.DB) comment.CommentRepository {
	return &CommentGormRepo{conn: db}
}

//comments returns all user comments stored in the database
func (commtRepo *CommentGormRepo) Comments() ([]models.Comment, []error) {
	ans := []models.Comment{}
	errs := commtRepo.conn.Find(&ans).GetErrors()
	if len(errs) > 0 {
		
		return nil, errs
	}
	return ans, errs
}
