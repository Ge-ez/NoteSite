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
//GetComments (for user)handles GET /comments request
func (ah *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	comments, errs := ah.commentService.Comments()

	if len(errs) > 0 {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comments, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "applcation/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}

//Getcomment handles GET comments/:id request
func (ah *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
        	if err != nil {
        		w.Header().Set("Content-Type", "application/json")
        		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
        		return
        	}

	comment, errs := ah.commentService.Comment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comment, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
