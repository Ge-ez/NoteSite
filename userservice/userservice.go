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
func Signin(username string, repo user.UserRepository) (*models.Usercheck, error) {
  login, err := repo.Login(username)
  if err != nil {
    return nil, err
  }
  return login, nil

}
