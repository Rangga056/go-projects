package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{} //This is a slice

	fmt.Printf("Welcome to our %v application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v still avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint //uint = postifive number
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter your number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry there are only %v tickets available\n", remainingTickets)
		} else {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conirmation email at at %v\n", firstName, lastName, userTickets, email)
		}

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			//NOTE:_ are called blank identifier to ignore a variable u dont want to use

			// iterates through all the array element and save the firstName to a slice
			var names = strings.Fields(booking) //strings.Fields() to split the string
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0 // check if there are no more tickets
		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked out. Come back next time.")
			break
		}
	}
}
