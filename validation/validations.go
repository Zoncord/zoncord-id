package validation

func Validate(value string, validation []func(string) error) error {
	for _, function := range validation {
		output := function(value)
		if output != nil {
			return output
		}
	}
	return nil
}

var SimpleValidation = []func(string) error{
	valueExistRule,
	spacesForbiddenRule,
}
