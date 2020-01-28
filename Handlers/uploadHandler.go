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

func (upldhand *NewUploadHandler) UploadNoteHandler(w http.ResponseWriter, r *http.Request) {
exts := []string{".pdf", ".docx", ".sys", ".txt", ".doc", ".ppt" , ".pptx",".zip"}
	token, err := rtoken.CSRFToken(upldhand.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		uploadForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		upldhand.tmpl.ExecuteTemplate(w, "uploadNote.layout", uploadForm)
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		uploadForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		uploadForm.Required("title", "description")
		uploadForm.MinLength("title",5)
		uploadForm.MinLength("description", 10)
		uploadForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !uploadForm.Valid() {
			upldhand.tmpl.ExecuteTemplate(w, "uploadNote.layout", uploadForm)
			return
		}
		mf, fh, err := r.FormFile("upload")
		if err != nil {
			uploadForm.VErrors.Add("upload", "File error")
			upldhand.tmpl.ExecuteTemplate(w, "uploadNote.layout", uploadForm)
			return
		}
		if !form.ExtensionChecker(fh.Filename,exts){
		    uploadForm.VErrors.Add("upload", "You can only upload the listed File Types!")
            upldhand.tmpl.ExecuteTemplate(w, "uploadNote.layout", uploadForm)
   			return
		}
		defer mf.Close()
		upld := &models.Upload{
			File_Name:        r.FormValue("title"),
			File_Description: r.FormValue("description"),
			File: fh.Filename,
			File_Uploader: upldhand.LoggedInUser.ID,
			File_Uploaded_to : upldhand.LoggedInUser.Course,
			File_Type: FilePath.Ext(fh.Filename),
			File:Uploaded_On: time.Now(),
		}
		UploadFile(&mf, fh.Filename)
		_, errs := upldhand.uploadService.StoreNotes(upld)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/homepage", http.StatusSeeOther)
	}
}

func (upldhand *UploadHandler) DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		id = uint(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := upldhand.uploadService.DeleteNote(id)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/mynotes", http.StatusSeeOther)
}
