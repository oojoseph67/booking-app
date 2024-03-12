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

// go run .
// go run main.go

package main // using the default package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var conferenceName = "mcQu33n GO Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}

	// range iterates over elements for different data structures (so not only arrays and slices)
	// underscores (_) are used as blank identifiers
	for _, bookingElement := range bookings {
		// var names = strings.Fields(bookingElement)
		firstNames = append(firstNames, bookingElement.firstName)
	}

	return firstNames
	// fmt.Printf("these are all our bookings: %v\n", bookings)
	// fmt.Printf("these are all the first names: %v\n", firstNames)
	// fmt.Println(" ")
}


func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
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

	return firstName, lastName, email, userTickets

}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {

	remainingTickets = remainingTickets - userTickets

	// creating a map for a user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("thank you %v %v for booking %v tickets, you will receive a confirmation email @ %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// fmt.Printf("slice length: %v\n", len(bookings))
}

func sendTIcket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets fro %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}

func main() {

	greetUsers()

	// for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := helper.Validation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {

			bookTicket(firstName, lastName, userTickets, email)
			fmt.Println(" ")
			wg.Add(1)
			go sendTIcket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("these are all the first names: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program

				fmt.Println("All tickets have been sold")
				fmt.Println("our conference is booked out. come back next year.")
				// break
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

			fmt.Println(" ")
			// break
			// continue
		}
	// }
	wg.Wait()
}
