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

//UpdatePost updates a given Post
func (ps *PostService) UpdatePost(post *models.Post) (*models.Post, []error) {
	post, err := ps.postRep.UpdatePost(post)
	if len(err) > 0 {
		return nil, err
	}
	//DeletePost deletes a given Post
func (ps *PostService) DeletePost(id uint) (*models.Post, []error) {
	post, err := ps.postRep.DeletePost(id)
	if len(err) > 0 {
		return nil, err
	}
	return post, err
}
	return post, err
}
