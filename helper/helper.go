// to export a function in GO we just capitalize the first letter of the function

package helper

import "strings"

func Validation(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}