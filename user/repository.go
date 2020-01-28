package user

import (
	"github.com/notesite/models"
)


//UserRepository repository(interface) specifies User user related database operations
type UserRepository interface {
    Users() ([]models.User,[]error)
    GetUser(id uint) (*models.User, []error)
    GetUserByUsername(username string) (*models.User,[]error)
	RegisterUser(user *models.User)(*models.User,[]error)
	EditUser(user *models.User)(*models.User,[]error)
	DeleteUser(id uint)(*models.User,[]error)
	UsernameExists(username string) bool

}


// SessionRepository specifies logged in user session related database operations
type SessionRepository interface {
	Session(sessionID string) (*models.Session, []error)
	StoreSession(session *models.Session) (*models.Session, []error)
	DeleteSession(sessionID string) (*models.Session, []error)
}
