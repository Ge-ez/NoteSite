package repository

import (
	"github.com/notesite/models"
	"github.com/notesite/posts"
	"github.com/jinzhu/gorm"
)

//PostGormRepo implements Post.PostRepository interface
type PostGormRepo struct {
	conn *gorm.DB
}

//NewPostGormRepo returns new object of PostGormRepo
func NewPostGormRepo(db *gorm.DB) posts.PostRepository {
	return &PostGormRepo{conn: db}
}

//Posts returns all user Posts stored in the database
func (postRepo *PostGormRepo) Posts() ([]models.Post, []error) {
	posts := []models.Post{}
	errs := postRepo.conn.Find(&posts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return posts, errs
}
