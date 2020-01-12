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
