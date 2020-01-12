package posts

import "github.com/notesite/models"


//PostService represents user - post operations
type PostService interface {
	Posts() ([]models.Post, []error) // gets array of Post struct or returns array of errors
	GetPost(id uint) (*models.Post, []error)
	UpdatePost(post *models.Post) (*models.Post, []error)
	DeletePost(id uint) (*models.Post, []error)
	StorePost(post *models.Post) (*models.Post, []error)
}
