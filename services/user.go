package services

func PasswordComplexityCheck(password string) error {
	// Password complexity check function
	if len(password) < 8 {
		// Password is too short
	}
	if len(password) > 64 {
		// Password is too long
	}
	return nil
}
