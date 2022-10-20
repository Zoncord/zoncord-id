package models

import "fmt"

var (
	InvalidEmailOrPassword = fmt.Errorf("invalid login or password")
	ErrInvalidCredentials  = fmt.Errorf("invalid credentials")
	ErrDatabaseNotAvaible  = fmt.Errorf("database not avaible")
	ErrInternalServerError = fmt.Errorf("internal server error")
	ErrInvalidGrant        = fmt.Errorf("invalid grant")
	ErrInvalidRedirectUri  = fmt.Errorf("invalid redirect_uri")
	DatabaseNotAvailable   = fmt.Errorf("database is not available")
	InvalidToken           = fmt.Errorf("invalid token")
)
