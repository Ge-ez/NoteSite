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
