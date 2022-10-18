package validation

import (
	"go.uber.org/zap"
	"regexp"
)

type validationValue struct {
	// Class to validate value
	title, value string
}

type twoValidationValues struct {
	// Class to validate two values
	value1, value2 validationValue
}

func newPasswordValidationValue(value string) validationValue {
	return validationValue{
		"password",
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
	// Function that checks if validation rules pass
	for _, rule := range validateRules {
		output := rule()
		if output != nil {
			return output
		}
	}
	return nil
}

func SimpleValidationRules(v validationValue) []func() error {
	// Some validation rules to validate simple strings (names, e.t.c)
	return []func() error{
		v.valueExistRule,
		v.spacesForbiddenRule,
		v.specialCharactersForbiddenRule,
		v.numbersForbiddenRule,
		v.maxValueLengthRule,
	}
}

func PasswordValidationRules(v validationValue) []func() error {
	// Some validation rules to validate password
	return []func() error{
		v.valueExistRule,
		v.spacesForbiddenRule,
		v.minPasswordLengthRule,
		v.maxValueLengthRule,
		v.numbersRequiredRule,
	}
}
func PasswordsValidationRules(v twoValidationValues) []func() error {
	// Validation rule to validate two passwords
	return []func() error{
		v.equivalencyRule,
	}
}

func PasswordValidation(password string) error {
	// Function that checks all password rules
	zap.L().Info("started password validation")
	val := newPasswordValidationValue(password)
	err := CheckRules(PasswordValidationRules(val))
	if err != nil {
		zap.L().Info(err.Error())
		return err
	}
	zap.L().Info("finished password validation")
	return nil
}

func PasswordsValidation(password1 string, password2 string) error {
	// Function that checks all passwords rules
	zap.L().Info("started passwords validation")
	passwords := twoValidationValues{
		newPasswordValidationValue(password1),
		newPasswordValidationValue(password2),
	}
	err := CheckRules(PasswordsValidationRules(passwords))
	if err != nil {
		zap.L().Info(err.Error())
		return err
	}
	// after checking PasswordsValidationRules two passwords are equal
	err = PasswordValidation(password1)
	if err != nil {
		zap.L().Info(err.Error())
		return err
	}
	zap.L().Info("finished passwords validation")
	return nil
}

func SimpleValidation(title, value string) error {
	// Function that checks all simple rules to passed value
	zap.L().Info("started " + title + " validation")
	var val = validationValue{
		title,
		value,
	}
	err := CheckRules(SimpleValidationRules(val))
	if err != nil {
		zap.L().Info(err.Error())
		return err
	}
	zap.L().Info("finished " + title + " validation")
	return nil
}

// EmailValidation TODO: remake email validation to array fo rules
func EmailValidation(email string) error {
	zap.L().Info("started email validation")
	emailRegex := regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	if !emailRegex.MatchString(email) {
		zap.L().Info(ErrInvalidEmailFormat.Error())
		return ErrInvalidEmailFormat
	}
	zap.L().Info("finished email validation")
	return nil
}
