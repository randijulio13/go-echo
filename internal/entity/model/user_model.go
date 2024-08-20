package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"ID" gorm:"not null;primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email,omitempty" gorm:"default:null"`
	gorm.Model
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
}
