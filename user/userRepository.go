package user

import (
  "github.com/notesite/models"
)

type UserRepository interface {
  Save(*models.User) error
  Login(username string) (*models.Usercheck, error)
  UpdateUser(id string,*models.User) error

}
 
