package validation

import "fmt"

func IsEq(err1 error, err2 error) bool {
	// comparing if errors are equal
	if err1 == err2 {
		// err1 and err2 are nil
		return true
	}
	if (err1 != nil && err2 != nil) && err1.Error() == err2.Error() {
		return true
	}
	return false
}

var (
	ErrPasswordsDontMatch = fmt.Errorf("passwords don't match")
	ErrInvalidEmailFormat = fmt.Errorf("invalid email format")
)

func (v *validationValue) ErrValueMustIncludeNumber() error {
	return fmt.Errorf("%s must include number", v.title)
}

func (v *validationValue) ErrValueTooShort() error {
	return fmt.Errorf("%s is too short", v.title)
}
func (v *validationValue) ErrValueTooLong() error {
	return fmt.Errorf("%s is too long", v.title)
}

func (v *validationValue) ErrRequiredValue() error {
	return fmt.Errorf("%s is required", v.title)
}

func (v *validationValue) ErrSpacesForbidden() error {
	return fmt.Errorf("spaces are forbidden in %s", v.title)
}

func (v *validationValue) ErrNumbersForbidden() error {
	return fmt.Errorf("numbers are forbidden in %s", v.title)
}

func (v *validationValue) ErrSpecialCharactersForbidden() error {
	return fmt.Errorf("special characters are forbidden in %s", v.title)
}
