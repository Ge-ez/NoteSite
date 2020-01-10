package user

import (
  "github.com/notesite/models"
)

type UserRepository interface {
  Save(*models.User) error
}
