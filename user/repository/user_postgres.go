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


