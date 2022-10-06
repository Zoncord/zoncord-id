package models

import (
	"github.com/Zoncord/zoncord-id/errors"
	"github.com/Zoncord/zoncord-id/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	// oauth2 info
	Applications  []Application  `json:"applications"`
	Grants        []Grant        `json:"grants"`
	AccessTokens  []AccessToken  `json:"access_token"`
	RefreshTokens []RefreshToken `json:"refresh_token"`
}

func CheckAuth(email string, password string) error {
	var user User
	err := db.First(&user, "email = ?", email).Error
	if err == gorm.ErrRecordNotFound {
		return errors.InvalidEmailOrPassword
	}

	if err != nil {
		return errors.DatabaseNotAvailable
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.InvalidEmailOrPassword
	}

	return nil
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
	user.FirstName = "123"
	db.Save(&user)
	return nil
}
