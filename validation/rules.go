package validation

import "go.uber.org/zap"

const MinPasswordLength int = 10
const MaxStringLength int = 64

func (v *validationValue) minPasswordLengthRule() error {
	zap.L().Info("checking minPasswordLengthRule")
	if len([]rune(v.value)) < MinPasswordLength {
		zap.L().Info(v.ErrValueTooShort().Error())
		return v.ErrValueTooShort()
	}
	zap.L().Info("finished minPasswordLengthRule")
	return nil
}

func (v *validationValue) maxValueLengthRule() error {
	zap.L().Info("checking maxValueLengthRule")
	if len([]rune(v.value)) > MaxStringLength {
		zap.L().Info(v.ErrValueTooLong().Error())
		return v.ErrValueTooLong()
	}
	zap.L().Info("finished maxValueLengthRule")
	return nil
}

func (v *validationValue) valueExistRule() error {
	zap.L().Info("checking valueExistRule")
	if len([]rune(v.value)) == 0 {
		zap.L().Info(v.ErrRequiredValue().Error())
		return v.ErrRequiredValue()
	}
	zap.L().Info("finished valueExistRule")
	return nil
}

func (v *validationValue) spacesForbiddenRule() error {
	zap.L().Info("checking spacesForbiddenRule")
	for _, character := range v.value {
		if character == ' ' || character == '\t' || character == '\n' {
			zap.L().Info(v.ErrSpacesForbidden().Error())
			return v.ErrSpacesForbidden()
		}
	}
	zap.L().Info("finished spacesForbiddenRule")
	return nil
}
func (v *validationValue) numbersForbiddenRule() error {
	zap.L().Info("Checking numbersForbiddenRule")
	for _, letter := range v.value {
		if '0' <= letter && letter <= '9' {
			zap.L().Info(v.ErrNumbersForbidden().Error())
			return v.ErrNumbersForbidden()
		}
	}
	zap.L().Info("finished numbersForbiddenRule")
	return nil
}

func (v *validationValue) numbersRequiredRule() error {
	zap.L().Info("Checking numbersRequiredRule")
	for _, letter := range v.value {
		if '0' <= letter && letter <= '9' {
			zap.L().Info("finished numbersRequiredRule")
			return nil
		}
	}
	zap.L().Info(v.ErrValueMustIncludeNumber().Error())
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
			zap.L().Info(v.ErrSpecialCharactersForbidden().Error())
			return v.ErrSpecialCharactersForbidden()
		}
	}
	zap.L().Info("finished specialCharactersForbiddenRule")
	return nil
}

func (v *twoValidationValues) equivalencyRule() error {
	zap.L().Info("Checking equivalencyRule")
	if v.value1.value != v.value2.value {
		// TODO: make formatting error message
		zap.L().Info(ErrPasswordsDontMatch.Error())
		return ErrPasswordsDontMatch
	}
	zap.L().Info("finished equivalencyRule")
	return nil
}
