package models

import "fmt"

var (
	InvalidEmailOrPassword = fmt.Errorf("invalid login or password")
	DatabaseNotAvailable   = fmt.Errorf("database is not available")
	InvalidToken           = fmt.Errorf("invalid token")
)
