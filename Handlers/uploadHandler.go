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
