package models

import (
	"ZoncordID/services"
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

func CreateUser(email string, password string, firstName string, lastName string) error {
	// create user
	var user User
	// set information about user
	user.Email = email
	user.Password = services.PasswordHasher(password)
	user.FirstName = firstName
	user.LastName = lastName
	user.IsActive = true
	user.IsSuperUser = false

	db.Create(&user)

	return nil
}
