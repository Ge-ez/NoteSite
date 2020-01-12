package repository

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"errors"

	"github.com/notesite/models"
)

type UserRepositoryImpl struct {
	conn *gorm.DB
}

func NewUserRepositoryImpl(Conn *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: Conn}
}


func (uri *UserRepositoryImpl) RegisterUser(user *models.User)(*models.User ,error) {
	users := user
//username taken?
_,err := uri.conn.Raw("SELECT * FROM users WHERE username = ?",user.Username).Rows()
if err != nil{
	return nil,errors.New("user name already taken try other")
}
errs := uri.conn.Create(users).GetErrors()
if len(errs) > 0 {
	return nil, errors.New("selection has failed")
}
return users,nil

	// //username taken?
	// _,err := uri.conn.Query("SELECT * FROM users where username = $1", user.UserName)
	// if err != nil {
	// 	//fmt.Println("here is the problem")
	// 	//panic(err)
	// 	return errors.New("user name already taken try other")
	// }
	// _, err = uri.conn.Exec("INSERT INTO users (username,first_name,last_name,email,password,phone,image) VALUES ($1,$2,$3,$4,$5,$6,$7)", user.UserName, user.FirstName,user.LastName, user.Email,user.Password, user.Phone, user.Image)

	// if err != nil {
	// 	fmt.Println("check me")
	// 	return errors.New("insertion has failed")
	// }
	// return nil

}
