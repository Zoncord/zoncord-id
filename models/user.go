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

func CheckAuth(email string, password string) (User, error) {
	var user User
	err := db.First(&user, "email = ?", email).Error
	if err == gorm.ErrRecordNotFound {
		return user, errors.InvalidEmailOrPassword
	}

	if err != nil {
		return user, errors.DatabaseNotAvailable
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.InvalidEmailOrPassword
	}

	return user, nil
}

func CreateUser(email string, password string, firstName string, lastName string) (User, error) {
	var user User
	user.Email = email
	user.Password = services.PasswordHasher(password)
	user.FirstName = firstName
	user.LastName = lastName
	user.IsActive = true
	user.IsSuperUser = false
	db.Create(&user)
	db.Save(&user)
	return user, nil
}

func GetUserByToken(token string) (User, error) {
	var user User
	var accessToken AccessToken
	err := db.First(&accessToken, "token = ?", token).Error
	if err == gorm.ErrRecordNotFound {
		return user, errors.InvalidToken
	}
	if err != nil {
		return user, errors.DatabaseNotAvailable
	}
	err = db.First(&user, "id = ?", accessToken.UserID).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
