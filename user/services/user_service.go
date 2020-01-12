package services

import (
	"github.com/notesite/entity"
	"github.com/notesite/user"
)

type UserServiceImpl struct {
	userRepo user.UserRepository
}

func NewUserServiceImpl(UserRep user.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: UserRep}
}

func (usi *UserServiceImpl) RegisterUser(user *models.User) (*models.User,error) {
	usr,err := usi.userRepo.RegisterUser(user)
	if err != nil {
		return usr,err
	}
	return usr,nil
}
func (usi *UserServiceImpl) AuthenticateUser(userName string, password string) (*models.Usercheck, error) {
	user, err := usi.userRepo.AuthenticateUser(userName, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (usi *UserServiceImpl) GetUser(userName string) (*models.User, error) {
	//check username?
	user, err := usi.userRepo.GetUser(userName)
	if err != nil {
		return user, err
	}
	return user, nil

}
