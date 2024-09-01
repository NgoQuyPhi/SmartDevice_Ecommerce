package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginData struct {
	Username string `json:"username" gorm:"unique"`
	Pass     string `json:"password"`
}
type User struct {
	gorm.Model
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	LoginData
	Email string `json:"email" gorm:"unique"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Pass = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) (bool, error) {
	if user.Pass == providedPassword {
		return true, nil
	} else {
		return false, errors.New("wrong password")
	}
}
