package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is  %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking applicattion\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickerts here to attend")

	var bookings []string

	for {
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

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so yo can't book %v tickets\n", remainingTickets, userTickets)
			break
		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first array: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking) 
			// var firstName = names[0]
			firstNames = append(firstNames, names[0])
		}


		fmt.Printf("The first names of our bookings are: %v\n", firstNames)
		if remainingTickets == 0 {
			//end the program
			fmt.Printf("Our conference is booked out, come back next year")
			break
		} 
	}
}

  
