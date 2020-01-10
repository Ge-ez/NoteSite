package user

import (
  "fmt"
  "database/sql"
  "github.com/notesite/models"
  "time"
)

type UserRepositoryPostgres struct {
  db *sql.DB
}

func NewUser() *models.User {
  return &models.User{
    Joindate: time.Now(),
  }
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
  return &UserRepositoryPostgres{db}
}

func (user *UserRepositoryPostgres) Save(signup *models.User) error {

  _, err := user.db.Exec("INSERT INTO users(name,username,email,password,gender,role,course,token,joindate) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)",signup.Name, signup.Username, signup.Email, signup.Password, signup.Gender, signup.Role, signup.Course, "", signup.Joindate)

  if err != nil {
    return err
    fmt.Println("unable to insert the data")
  }
  return nil

}
