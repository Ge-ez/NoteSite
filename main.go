package main

import (
  "html/template"
  "net/http"
  "github.com/julienschmidt/httprouter"
	"github.com/notesite/handlers"
	postrepo "github.com/notesite/posts/repository"
	postserv "github.com/notesite/posts/service"
	comntserv "github.com/notesite/comment/service"
	comntrepo "github.com/notesite/comment/repository"
	"github.com/notesite/database"

)

var tmpl *template.Template

func init() {
  tmpl = template.Must(template.ParseGlob("Views/temp_user/*"))
}

func main() {
 
}

func index(w http.ResponseWriter, r *http.Request) {
  tmpl.ExecuteTemplate(w, "index.html", nil)
}
