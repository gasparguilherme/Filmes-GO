package validate

import "errors"

func ValidateID(id int) error {
	if id <= 0 {
		return errors.New("please enter an id greater than zero")
	}
	return nil

}
