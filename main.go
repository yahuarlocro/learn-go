package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is  %T\n", conferenceTickets, remainingTickets, conferenceName)
	// fmt.Printf("Welcome to our %v booking applicattion\n", conferenceName)

	// for remainingTickets > 0 && len(bookings) < 50 {
	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// if userTickets <= remainingTickets {
			bookTicket(userTickets, firstName, lastName, email)

			var firstNames = getFirstNames()
			fmt.Printf("The first names of our bookings are: %v\n", firstNames)

			if remainingTickets == 0 {

				//end the program
				fmt.Printf("Our conference is booked out, come back next year")
				break
			}
			// } else if userTickets == remainingTickets {
			// 	// do something else
		} else {
			if !isValidName {
				fmt.Printf("First name or last name is too short\n")
			}

			if !isValidEmail {
				fmt.Printf("email addres you entered does not contain a @\n")

			}

			if !isValidTicketNumber {
				fmt.Printf("number of tickets you entered is invalid\n")
			}

			// fmt.Printf("We only have %v tickets remaining, so yo can't book %v tickets\n", remainingTickets, userTickets)
			// fmt.Printf("Your input data is invalid, please try again\n")
			// continue
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickerts here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	isValidEmail := strings.Contains(email, "@")

	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	// userName = "Tom"

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first array: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
