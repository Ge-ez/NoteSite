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


}
func (usi *UserServiceImpl) GetUser(userName string) (*models.User, error) {
	//check username?
	user, err := usi.userRepo.GetUser(userName)
	if err != nil {
		return user, err
	}
	return user, nil

}
func (usi *UserServiceImpl) EditUser(user *models.User)(*models.User,[]error) {
	urs,err := usi.userRepo.EditUser(user)
	if err != nil {
		return urs,err
	}
	return urs,nil
}
func (usi *UserServiceImpl) DeleteUser(id uint)(*models.User,error) {
	urs,err := usi.userRepo.DeleteUser(id)
	if err != nil {
		return urs,nil
	}
	return urs,nil
}
