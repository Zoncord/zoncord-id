package errors

import "fmt"

var (
	NotFoundError = fmt.Errorf("resource could not be found")
)
