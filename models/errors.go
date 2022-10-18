package models

import "fmt"

var (
	// TODO: move to another files
	InvalidEmailOrPassword = fmt.Errorf("invalid login or password")
	DatabaseNotAvailable   = fmt.Errorf("database is not available")
	InvalidToken           = fmt.Errorf("invalid token")
)
