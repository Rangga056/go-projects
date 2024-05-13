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

	greetUser(conferenceName, int(conferenceTickets), uint(remainingTickets))

	for remainingTickets > 0 && len(bookings) < 50 {
		//NOTE:You can also put condition in for loop to keep iterating as long as condition is true
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

		//Validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") //Check if the email input contain "@"
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isValidCity := city == "Singapore" || city == "London"
		// Or operator

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conirmation email at at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// call function print firstName
			printFirstNames(bookings)

			noTicketsRemaining := remainingTickets == 0 // check if there are no more tickets
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next time.")
				break
			}
		} else {
			fmt.Println("Your input data is invalid, please try again")
			// fmt.Printf("Sorry there are only %v tickets available\n", remainingTickets)
		}
	}
}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		//NOTE:_ are called blank identifier to ignore a variable u dont want to use

		// iterates through all the array element and save the firstName to a slice
		var names = strings.Fields(booking) //strings.Fields() to split the string
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These are all our bookings: %v\n", firstNames)
}
