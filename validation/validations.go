package validation

import (
	"go.uber.org/zap"
	"regexp"
)

type validationValue struct {
	title, value string
}

type twoValidationValues struct {
	value1, value2 validationValue
}

func newPasswordValidationValue(value string) validationValue {
	return validationValue{
		"password",
		value,
	}
}

func newFirstNameValidationValue(value string) validationValue {
	return validationValue{
		"first name",
		value,
	}
}
func newLastNameValidationValue(value string) validationValue {
	return validationValue{
		"last name",
		value,
	}
}
func newTestValidationValue(value string) validationValue {
	return validationValue{
		"test value",
		value,
	}
}

func CheckRules(validateRules []func() error) error {
	for _, rule := range validateRules {
		output := rule()
		if output != nil {
			return output
		}
	}
	return nil
}

func SimpleValidationRules(v validationValue) []func() error {
	return []func() error{
		v.valueExistRule,
		v.spacesForbiddenRule,
		v.specialCharactersForbiddenRule,
		v.numbersForbiddenRule,
		v.maxValueLengthRule,
	}
}

func PasswordValidationRules(v validationValue) []func() error {
	return []func() error{
		v.valueExistRule,
		v.spacesForbiddenRule,
		v.minPasswordLengthRule,
		v.maxValueLengthRule,
		v.numbersRequiredRule,
	}
}
func PasswordsValidationRules(v twoValidationValues) []func() error {
	return []func() error{
		v.equivalencyRule,
	}
}

func PasswordsValidation(password1 string, password2 string) error {
	// Password validation function
	passwords := twoValidationValues{
		newPasswordValidationValue(password1),
		newPasswordValidationValue(password2),
	}
	err := CheckRules(PasswordsValidationRules(passwords))
	if err != nil {
		return err
	}

	err = CheckRules(PasswordValidationRules(passwords.value1))
	if err != nil {
		return err
	}
	return nil
}
func PasswordValidation(password string) error {
	val := newPasswordValidationValue(password)
	err := CheckRules(PasswordValidationRules(val))
	if err != nil {
		return err
	}
	return nil
}

func FirstNameValidation(value string) error {
	val := newFirstNameValidationValue(value)
	err := CheckRules(SimpleValidationRules(val))
	if err != nil {
		return err
	}
	return nil
}
func LastNameValidation(value string) error {
	val := newLastNameValidationValue(value)
	err := CheckRules(SimpleValidationRules(val))
	if err != nil {
		return err
	}
	return nil
}

// EmailValidation TODO: remake email validation to array fo rules
func EmailValidation(email string) error {
	emailRegex := regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	if !emailRegex.MatchString(email) {
		zap.L().Info("Email address isn't walid")
		return ErrInvalidEmailFormat
	}
	zap.L().Info("Email address validated successfully")
	return nil
}
