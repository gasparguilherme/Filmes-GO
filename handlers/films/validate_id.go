package handlers

import "fmt"

func validateID(id int) error {
	if id <= 0 {
		fmt.Println("please enter an id greater than zero")
	}
	return nil

}
