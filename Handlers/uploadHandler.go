package handlers

import(
   "html/template"
    "net/http"
    "net/url"
    "time"
    "io"
   	"mime/multipart"
   	"os"
   	"path/filepath"
   	"strconv"

     "github.com/notesite/form"
     "github.com/notesite/models"
     "github.com/notesite/upload"
     )


type UploadHandler struct {
	tmpl           *template.Template
	uploadService    upload.UploadService
	loggedInUser   *models.User
	csrfSignKey []byte
}

func NewUploadHandler(t *template.Template,
    upldServ upload.UploadService,csKey []byte) *UploadHandler{
	return &UploadHandler{
	    tmpl: t,
	    uploadService: upldServ,
	    csrfSignKey: csKey}
}

func (upldhand *NewUploadHandler) UploadedNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, errs := ach.categorySrv.NotesByCourseName(upldhand.loggedInUser.Course)
	if errs != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Notes []entity.Category
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Notes: notes,
		CSRF:       token,
	}
	upldhand.tmpl.ExecuteTemplate(w, "homepage.layout", tmplData)
}
