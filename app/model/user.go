package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	idNumber string 
	fullName string
	sex string
	bod time.Time // birthday
	pod string // place of birth
	room string
}