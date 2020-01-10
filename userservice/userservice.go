package userservice

import (
  "github.com/notesite/models"
  "github.com/notesite/user"
)

func Signup(user *models.User, repo user.UserRepository) error {
  err := repo.Save(user)
  if err != nil {
    return err
  }
  return nil
}
