package models

import (
	"ZoncordID/services"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var db = GetDB()

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	// personal info
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	// auth info
	IsActive    bool `json:"is_active"`
	IsSuperUser bool `json:"is_superuser"`
}

func CheckAuth(email string, password string) (bool, error) {
	var user User
	err := db.First(&user, "email = ?", email).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CreateUser(email string, password string, firstName string, lastName string) error {
	var user User
	user.Email = email
	user.Password = services.PasswordHasher(password)
	user.FirstName = firstName
	user.LastName = lastName
	user.IsActive = true
	user.IsSuperUser = false
	db.Create(&user)
	return nil
}
