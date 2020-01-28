package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"

	"github.com/notesite/database"
	"github.com/notesite/handlers"
	"github.com/notesite/models"

	tkn "github.com/notesite/security/token"
	postrepo "github.com/notesite/posts/repository"
	postserv "github.com/notesite/posts/service"

	comntserv "github.com/notesite/comment/service"
	comntrepo "github.com/notesite/comment/repository"


	upldServ "github.com/notesite/upload/service"
	upldRepo "github.com/notesite/upload/repository"

	userrepo "github.com/notesite/user/repository"
	userserv "github.com/notesite/user/services"

)
var tmpl *template.Template

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&models.Session{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}
func main() {

	csrfSignKey := []byte(tkn.GenerateRandomID(32))
	tmpl := template.Must(template.ParseGlob("views/*.html"))

	db:=database.Conn()
	defer db.Close()

	errs := createTables(db)
	if len(errs)>0{
		fmt.Println(errs)
	}
    sessionRepo := userrepo.NewSessionGormRepo(db)
    sessionSrv := userserv.NewSessionService(sessionRepo)

	userRepo := userrepo.NewUserRepositoryImpl(db)
	userServ := userserv.NewUserServiceImpl(userRepo)
	sess := configSess()
    uh := handlers.NewUserHandler(tmpl, userServ, sessionSrv, sess, csrfSignKey)

	comntRepo := comntrepo.NewCommentGormRepo(db)
	comntServ := comntserv.NewCommentService(comntRepo)
	commentHandler := handlers.NewCommentHandler(comntServ)

	postRepo := postrepo.NewPostGormRepo(db)
	postService := postserv.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	upldrepo := upldRepo.NewUploadRepositoryImpl(db)
	upldserv := upldServ.NewUploadServiceImpl(upldrepo)
	uploadhandler := handlers.NewUploadHandler(tmpl,upldserv,csrfSignKey)

	router := httprouter.New()

	router.GET("/posts", postHandler.GetPosts)
	router.GET("/posts/:id", postHandler.GetPost)
	router.POST("/post", postHandler.Postposts)
	router.PUT("/posts/:id", postHandler.PutPost)
	router.DELETE("/posts/:id", postHandler.DeletePost)

	router.GET("/comments", commentHandler.GetComments)
	router.GET("/comments/:id", commentHandler.GetComment)
	router.POST("/comment", commentHandler.PostComment)
	router.PUT("/comments/:id", commentHandler.PutComment)
	router.DELETE("/comments/:id", commentHandler.DeleteComment)

    fs := http.FileServer(http.Dir("views/assets"))
    http.Handle("/assets/", http.StripPrefix("/assets/", fs))
    http.HandleFunc("/", uh.Login)
   	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
   	http.HandleFunc("/signup", uh.Signup)
	//http.HandleFunc("/posts", postHandler.FetchPosts)
	//http.HandleFunc("/posts/post", postHandler.StorePosts)
	//http.HandleFunc("/homepage",index)
    http.HandleFunc("/homepage",uploadhandler.UploadedNotesHandler)
    http.HandleFunc("/upload",uploadhandler.UploadNoteHandler)
	http.ListenAndServe(":8080", nil)

}

func configSess() *models.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := tkn.GenerateRandomID(32)
	signingString, err := tkn.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &models.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}

func index(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w,"aa..html",nil)
}
