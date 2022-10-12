package errors

import "fmt"

var (
	NotFoundError          = fmt.Errorf("resource could not be found")
	UserNotRegistered      = fmt.Errorf("user is not registered")
	InvalidEmailOrPassword = fmt.Errorf("invalid login or password")
	DatabaseNotAvailable   = fmt.Errorf("database is not available")
	InvalidToken           = fmt.Errorf("invalid token")
	PasswordTooShort       = fmt.Errorf("password is too short")
	PasswordTooLong        = fmt.Errorf("password is too long")
)
