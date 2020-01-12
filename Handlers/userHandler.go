package handlers


import (
	"github.com/notesite/models"
	"github.com/notesite/user"
	"html/template"

    "time"
	"golang.org/x/crypto/bcrypt"
	
	"net/http"

	uuid "github.com/satori/go.uuid"

)

//UserHandler handles user related requests
type UserHandler struct {
	tmpl    *template.Template
	userSrv user.UserService
}


var dbSessions = map[string]string{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(T *template.Template, US user.UserService) *UserHandler {
	return &UserHandler{tmpl: T, userSrv: US}
}
