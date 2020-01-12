package user

import (
	"github.com/notesite/models"
)


//UserService ... this is our service usescase(has (interface)abstract classes that outer layers can use)
type UserService interface {
	RegisterUser(user *models.User)(*models.User,error)
	AuthenticateUser(userName string, password string) (*models.Usercheck, error)
	GetUser(userName string) (*models.User, error)
	EditUser(user *models.User)(*models.User,[]error)
	DeleteUser(id uint)(*models.User,error)
}
