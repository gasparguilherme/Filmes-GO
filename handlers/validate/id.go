package validate

import "fmt"

func ValidateID(id int) error {
	if id <= 0 {
		fmt.Println("please enter an id greater than zero")
	}
	return nil

}
