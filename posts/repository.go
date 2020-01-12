package posts

import "github.com/notesite/models"


//PostRepository represents post related operations on the database
type PostRepository interface {
	Posts() ([]models.Post, []error)
	GetPost(id uint) (*models.Post, []error)
	UpdatePost(post *models.Post) (*models.Post, []error)
	DeletePost(id uint) (*models.Post, []error)
	StorePost(post *models.Post) (*models.Post,[]error)
}
