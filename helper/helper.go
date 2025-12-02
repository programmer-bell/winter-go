package helper

import "strings"


func ValidateUserInput(fname string, lname string , email string , userTickets uint,reminingTickets uint ) (bool, bool, bool){
	// var isValidName bool = len(fname) >= 2 && len(lname) >= 2
	isValidName := len(fname) >= 2 && len(lname) >= 2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0  && userTickets <= reminingTickets
	return isValidName, isValidEmail,isValidTicketNumber
}




