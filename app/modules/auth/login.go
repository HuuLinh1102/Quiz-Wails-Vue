package auth

import (
	model "changeme/app/model"
	"errors"
)

// Login function to check if the user is in the database
type ILogin interface {
	Login(idnum string) (user model.User, err error)
}

type LoginService struct {
	// LoginRepo ILogin
}

func (ls *LoginService) Login(idnum string) (user model.User, err error) {
	if idnum == "itc" {
		user = model.User{}
	} else {
		err = errors.New("Không tìm thấy thí sinh trên hệ thống")
	}

	return
}
