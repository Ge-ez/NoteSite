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
func Findusername(username string, repo user.UserRepository) (*models.Usercheck, error) {
  login, err := repo.Login(username)
  if err != nil {
    return nil, err
  }
  return login, nil

}

func Update(user *models.User,repo UserRepository) error{
  err := repo.UpdateUser(user.ID,user)
  if err!=nil{
    return err
  }

  return nil
}
