package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	// personal info
	FirstName  string
	LastName   string
	MiddleName string
	Phone      string
	//
	IsActive    bool `json:"is_active"`
	IsSuperUser bool
	//
}

func CheckAuth(email string, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Email: email, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
