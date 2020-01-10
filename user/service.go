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
func (user *UserRepositoryPostgres) Login(username string) (*models.Usercheck, error) {
  query := SELECT username,password FROM users WHERE username = $1
  var login models.Usercheck
  statement, err := user.db.Prepare(query)
  if err != nil {
    return nil, err
  }
  defer statement.Close()

  err = statement.QueryRow(username).Scan(&login.Username, &login.Password)
  if err != nil {
    return nil, err
  }
  return &login, nil
}

func (user *UserRepositoryPostgres) UpdateUser(id string,update *models.User) error{
  query := "UPDATE users SET name=$1, username=$2,password=$3,email=$4,about=$5 WHERE id=$6"
  statement, err := user.db.Prepare(query)
  if err != nil {
    return err
  }
  defer statement.Close()
  _, err = statement.Exec(update.Name, update.Username, update.Email, update.Password, update.About)

  if err != nil {
    return err
  }
  return nil
}
