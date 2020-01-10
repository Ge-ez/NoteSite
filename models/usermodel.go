package models

import "time"

type User struct {
	ID       uint32
	Name     string
	Username string
	Email    string
	Password string
	Gender   string
	Role     string
	Course   string
	Image    string
	About    string
	Token    string
	Joindate time.Time
}

type Usercheck struct {
	Username, Password string
}
