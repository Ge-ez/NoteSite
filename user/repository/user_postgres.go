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

func (uri *UserRepositoryImpl) AuthenticateUser(userName string, password string) (*models.Usercheck, error) {
	user:= models.User{}

//is there a username?
rows,err := uri.conn.Raw("SELECT * FROM  users WHERE username = ?",userName).Rows()
defer rows.Close()

if (rows != nil){
	if err != nil{
		return &user,errors.New("username and/or password do not match")
	}
	for rows.Next(){
		uri.conn.ScanRows(rows,&user)
	}
		//does the entered password match with the stred password?
		err = bcrypt.CompareHashAndPassword(user.Password,[]byte(password))
		if err!= nil{
			return &user,errors.New("username and/or password do not match")
		}
		return &user,nil
}
return &user,errors.New("username and/or password do not match")


	// //is there a username?
	// row := uri.conn.QueryRow("SELECT * FROM users where username = $1", userName)
	
	// user := models.User{}
	// if row != nil {
	// 	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.UserName, &user.Email,&user.Password, &user.Phone, &user.Image)
	// 	if err != nil {
	// 		return user, errors.New("username  and/or password do not match")
	// 	}

	// 	//does the entered password match with the stred password?
	// err = bcrypt.CompareHashAndPassword(user.Password,[]byte(password))
	// if err!= nil{
	// 	return user,errors.New("username and/or password do not match")
	// }
		
	// 	return user, nil
	// }
	// return user, errors.New("username and/or password do not match")	
}
func (uri *UserRepositoryImpl) GetUser(userName string) (*models.User, error) {
	user:=models.User{}


	//check username if exist return users
	rows,err := uri.conn.Raw("SELECT * FROM users WHERE username = ?",userName).Rows()
	if rows != nil{
		for rows.Next(){
			uri.conn.ScanRows(rows,&user)
		}
		if err != nil{
			return &user,err
		}
		return &user,nil
	}
	return &user,errors.New("user not found")



	// // check username if exist return users
	// row := uri.conn.QueryRow("SELECT * FROM users where username = $1", userName)
	// user := models.User{}
	// if row != nil {
	// 	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.UserName, &user.Email,&user.Password, &user.Phone, &user.Image)
	// 	if err != nil {
	// 		return user, err
	// 	}
	// 	return user, nil
	// }
	// return user, errors.New("user not found")

}

//EditUser ... edit our user model
func (uri *UserRepositoryImpl) EditUser(user *models.User)(*models.User ,[]error) {
usr:= user
errs :=uri.conn.Save(usr).GetErrors()

if len(errs)>0{
	return nil,errs
}
return usr,nil

	// _, err := uri.conn.Exec("UPDATE users SET first_name = $1,last_name = $2,username = $3,email = $4,password= $5, phone = $6,image = $7 WHERE id = $8", user.FirstName, user.LastName, user.UserName, user.Email,user.Password, user.Phone, user.Image,user.UserID)
	// if err != nil {
	// 	return errors.New("Update has faild")
	// }
	// return nil
}

//DeleteUser ... Delete user
func (uri *UserRepositoryImpl) DeleteUser(id uint) (*models.User,error) {
	user := models.User{}
rows,err:= uri.conn.Raw("DELETE FROM users WHERE id = ?",id).Rows()
if rows != nil{
for rows.Next(){
	uri.conn.ScanRows(rows,&user)
}
if err != nil{
	return &user,err
}
return &user,nil
	
}
return &user,errors.New("user not found")


	// _, err := uri.conn.Exec("DELETE FROM users WHERE id = $1", id)
	// if err != nil {
	// 	return errors.New("Delete has faild")
	// }
	// return nil
	
	
}

