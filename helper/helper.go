package helper

import "strings"

// ! capitalizing the first letter of the function below to export it to be used by other
// ! packages

// ! NOTE: we can also capitalize the first letter of the  variable name to export it

func ValidateUserInput(firstName string, lastName string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	var isNameValid bool = len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isTicketNumberValid := tickets > 0 && tickets <= remainingTickets
	return isNameValid, isEmailValid, isTicketNumberValid
}
