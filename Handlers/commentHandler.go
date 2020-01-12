package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"github.com/notesite/models"
	"github.com/notesite/comment"
	"github.com/julienschmidt/httprouter"
)

//commentHandler handles menu related requests from a user
type CommentHandler struct {
	commentService comment.CommentService
}

//NewCommentHandler (for user) returns new CommentHandler object
func NewCommentHandler(cmntService comment.CommentService) *CommentHandler {
	return &CommentHandler{commentService: cmntService}
}
