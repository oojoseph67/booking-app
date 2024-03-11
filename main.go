// STEP 1: go mod init booking-app (project name)
// we are initializing the GO project
// STEP 2: NOTE: in GO out code must belong to a package
// example: package main
// STEP 3: NOTE: GO needs to know where to start executing the code
// (the code needs an entry point)

// NOTE: packages are a collection of source files that have specific functions
// NOTE: Variable names should be descriptive
// NOTE: GO is type strict and has a strict complier
// NOTE: in GO there is only one type of loop are FOR LOOP

package main // using the default package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "mcQu33n GO Conference"
	// with this type of declaration i cant assign the data type

	const conferenceTickets = 50
	var remainingTickets uint = 50

	var bookings []string // slice... the difference is var bookings [50]string

	// fmt.Printf("%T, %T, %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v conferenceName booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		var city string
		// ask user for their name

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println(" ")

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println(" ")

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println(" ")

		fmt.Println("Enter your required ticket: ")
		fmt.Scan(&userTickets)
		fmt.Println(" ")

		fmt.Println("Enter desired conference location")
		fmt.Scan(&city)
		fmt.Println(" ")

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets
		isValidCity := city != "Lagos" && city != "Abuja"

		if isValidName && isValidEmail && isValidTicket && isValidCity {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("thank you %v %v for booking %v tickets, you will receive a confirmation email @ %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			fmt.Printf("slice length: %v\n", len(bookings))

			firstNames := []string{}

			// range iterates over elements for different data structures (so not only arrays and slices)
			// underscores (_) are used as blank identifiers
			for _, bookingElement := range bookings {
				var names = strings.Fields(bookingElement)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("these are all our bookings: %v\n", bookings)
			fmt.Printf("these are all the first names: %v\n", firstNames)
			fmt.Println(" ")

			if remainingTickets == 0 {
				// end program

				fmt.Println("All tickets have been sold")
				fmt.Println("our conference is booked out. come back next year.")
				break
			}
		} else {
			fmt.Println(" ")

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email does contain the @ sign")
			}

			if !isValidTicket {
				fmt.Println("number of tickets you entered is invalid")
			}

			if !isValidCity {
				fmt.Println("city must be either Lagos or Abuja")
			}

			fmt.Println(" ")
			// break
			continue
		}
	}
}
