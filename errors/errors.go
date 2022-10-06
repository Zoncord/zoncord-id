package errors

import "fmt"

var (
	NotFoundError          = fmt.Errorf("resource could not be found")
	UserNotRegistered      = fmt.Errorf("user is not registered")
	InvalidEmailOrPassword = fmt.Errorf("invalid login or password")
	DatabaseNotAvailable   = fmt.Errorf("database is not available")
)
