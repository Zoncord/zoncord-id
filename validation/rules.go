package validation

import "go.uber.org/zap"

const MinPasswordLength int = 10
const MaxStringLength int = 64

func (v *validationValue) minPasswordLengthRule() error {
	zap.L().Info("Checking minPasswordLengthRule")
	if len([]rune(v.value)) < MinPasswordLength {
		return v.ErrValueTooShort()
	}
	return nil
}
func (v *validationValue) maxValueLengthRule() error {
	zap.L().Info("Checking maxValueLengthRule")
	if len([]rune(v.value)) > MaxStringLength {
		return v.ErrValueTooLong()
	}
	return nil
}

func (v *validationValue) valueExistRule() error {
	zap.L().Info("Checking valueExistRule")
	if len([]rune(v.value)) == 0 {
		return v.ErrRequiredValue()
	}
	return nil
}

func (v *validationValue) spacesForbiddenRule() error {
	zap.L().Info("Checking spacesForbiddenRule")
	for _, character := range v.value {
		if character == ' ' {
			return v.ErrSpacesForbidden()
		}
		if character == '\t' {
			// TODO make error handling
		}
		if character == '\n' {
			// TODO make error handling
		}
	}
	return nil
}
func (v *validationValue) numbersForbiddenRule() error {
	zap.L().Info("Checking numbersForbiddenRule")
	for _, letter := range v.value {
		if '0' <= letter && letter <= '9' {
			return v.ErrNumbersForbidden()
		}
	}
	return nil
}

func (v *validationValue) numbersRequiredRule() error {
	zap.L().Info("Checking numbersRequiredRule")
	for _, letter := range v.value {
		if '0' <= letter && letter <= '9' {
			return nil
		}
	}
	return v.ErrValueMustIncludeNumber()
}
func (v *validationValue) specialCharactersForbiddenRule() error {
	zap.L().Info("Checking specialCharactersForbiddenRule")
	for _, letter := range v.value {
		// checks !"#$%&'()*+,/:;<=>?[\]^ characters
		if ('!' <= letter && letter <= ',') ||
			'/' == letter ||
			(':' <= letter && letter <= '?') ||
			('[' <= letter && letter <= '^') {
			return v.ErrSpecialCharactersForbidden()
		}
	}
	return nil
}

func (v *twoValidationValues) equivalencyRule() error {
	zap.L().Info("Checking equivalencyRule")
	if v.value1.value != v.value2.value {
		// TODO: make formatting error message
		return ErrPasswordsDontMatch
	}
	return nil
}
