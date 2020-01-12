package repository
import (
	"github.com/notesite/models"
)


//UserRepository repository(interface) specifies User user related database operations
type UserRepository interface {
	RegisterUser(user *models.User)(*models.User,error)
	AuthenticateUser(userName string, password string) (*models.Usercheck, error)
	GetUser(userName string) (*models.User, error)
	EditUser(user *models.User)(*models.User,[]error)
	DeleteUser(id uint)(*models.User,error)
}
