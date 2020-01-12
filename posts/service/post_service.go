package service
import (
	"github.com/notesite/models"
	"github.com/notesite/posts"
)

//PostService implements post.Service
type PostService struct {
	postRep posts.PostRepository
}

// NewPostService returns a new PostService object
func NewPostService(postsRepo posts.PostRepository) posts.PostService {
	return &PostService{postRep: postsRepo}
}

//Posts returns all stored Posts
func (ps *PostService) Posts() ([]models.Post, []error) {
	posts, err := ps.postRep.Posts()
	if len(err) > 0 {
		return nil, err
	}
	return posts, err
}

//Post retrieves stored Post given its id
func (ps *PostService) GetPost(id uint) (*models.Post, []error) {
	post, err := ps.postRep.GetPost(id)
	if len(err) > 0 {
		return nil, err
	}
	return post, err
}
