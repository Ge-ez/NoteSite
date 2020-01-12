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
func (uh *UserHandler) alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, errr := uh.userSrv.GetUser(un)
	if errr != nil {
		return false
	}
	return true

}
func (uh *UserHandler) getUser(w http.ResponseWriter, req *http.Request) *models.User {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)

	//if the user exists already,get user
	var u *models.User
	if un, ok := dbSessions[c.Value]; ok {
		u, _ = uh.userSrv.GetUser(un)
	}
	return u
}

