package user

import (
  "github.com/notesite/models"
)

type UserRepository interface {
  Save(*models.User) error
  Findbyusername(username string) (*models.Usercheck, error)
  Update(id uint32,*models.User) error

}
 
