package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type SignupData struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
}

type LoginData struct {
	Username string `form:"taikhoan" json:"taikhoan" gorm:"unique"`
	Pass     string `form:"matkhau" json:"matkhau"`
}
type User struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) (bool, error) {
	if user.Password == providedPassword {
		return true, nil
	} else {
		return false, errors.New("wrong password")
	}
}
