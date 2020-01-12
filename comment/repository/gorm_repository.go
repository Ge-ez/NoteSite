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

//Comment returns a user Comment stored in the database which has the given id
func (commtRepo *CommentGormRepo) Comment(id uint) (*models.Comment, []error) {
	qstn := models.Comment{}
	errs := commtRepo.conn.Where("id = ?", id).First(&qstn).GetErrors()
	// errs := commtRepo.conn.First(&qstn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &qstn, errs
}

//UpdateComment updates a given Comment in the database
func (commtRepo *CommentGormRepo) UpdateComment(comment *models.Comment) (*models.Comment, []error) {
	commt := comment
	errs := commtRepo.conn.Save(commt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return commt, errs
}
