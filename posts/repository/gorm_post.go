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

//Post returns a user Post stored in the database which has the given id
func (postRepo *PostGormRepo) GetPost(id uint) (*models.Post, []error) {
	post := models.Post{}
	errs := postRepo.conn.Where("id = ?", id).First(&post).GetErrors()
	// errs := postRepo.conn.First(&post, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &post, errs
}

//UpdatePost updates a given Post in the database
func (postRepo *PostGormRepo) UpdatePost(posts *models.Post) (*models.Post, []error) {
	post := posts
	errs := postRepo.conn.Save(post).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return post, errs
}

//DeletePost deletes a Post with a given id from the database
func (postRepo *PostGormRepo) DeletePost(id uint) (*models.Post, []error) {
	post, errs := postRepo.GetPost(id)
	if len(errs) > 0 {
		return nil, errs
	}
	// errs := postRepo.conn.Where("id = ?", id).First(&post).GetErrors()
	errs = postRepo.conn.Delete(post).GetErrors()


	if len(errs) > 0 {
		return nil, errs
	}
	return post, errs
}
