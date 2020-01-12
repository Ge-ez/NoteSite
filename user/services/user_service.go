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
