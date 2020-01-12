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
