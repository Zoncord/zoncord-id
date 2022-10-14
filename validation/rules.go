package validation

import "github.com/Zoncord/zoncord-id/errors"

func valueExistRule(value string) error {
	if len([]rune(value)) == 0 {
		return errors.RequiredValue
	}
	return nil
}

func spacesForbiddenRule(value string) error {
	for _, character := range value {
		if character == ' ' {
			return errors.SpacesForbidden
		}
	}
	return nil
}
