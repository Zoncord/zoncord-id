package models

import (
	"github.com/Zoncord/zoncord-id/validation"
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
		return user, InvalidEmailOrPassword
	}

	if err != nil {
		return user, DatabaseNotAvailable
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, InvalidEmailOrPassword
	}

	return user, nil
}

func (u *User) Create(email string, password string, firstName string, lastName string) error {
	err := validation.EmailValidation(email)
	if err != nil {
		return err
	}
	u.Email = email

	err = validation.PasswordValidation(password)
	if err != nil {
		return err
	}
	u.Password = password

	err = validation.FirstNameValidation(firstName)
	if err != nil {
		return err
	}
	u.FirstName = firstName

	err = validation.LastNameValidation(lastName)
	if err != nil {
		return err
	}
	u.LastName = lastName
	u.IsActive = true
	u.IsSuperUser = false
	db.Create(&u)
	db.Save(&u)
	return nil
}

func GetUserByToken(token string) (User, error) {
	var user User
	var accessToken AccessToken
	err := db.First(&accessToken, "token = ?", token).Error
	if err == gorm.ErrRecordNotFound {
		return user, InvalidToken
	}
	if err != nil {
		return user, DatabaseNotAvailable
	}
	err = db.First(&user, "id = ?", accessToken.UserID).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
