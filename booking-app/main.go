package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) //This is a slice

// Create a struct in Go
// Struct  can have mixed data types
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// for remainingTickets > 0 && len(bookings) < 50 {
	//NOTE:You can also put condition in for loop to keep iterating as long as condition is true
	firstName, lastName, email, userTickets := getUserInput()

	//Validate user input
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1) //NOTE:Add number of goroutines to wait for (number based on the running goroutines)
		go sendTicket(userTickets, firstName, lastName, email)
		// call function print firstName
		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0 // check if there are no more tickets
		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked out. Come back next time.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you enter is too short")
		}
		if !isValidEmail {
			fmt.Println("your email address doesn't contain a @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait() //NOTE:Block until the waitgroup counter is 0
	// }
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//NOTE:_ are called blank identifier to ignore a variable u dont want to use

		// iterates through all the array element and save the firstName to a slice
		// var names = strings.Fields(booking) //strings.Fields() to split the string
		firstNames = append(firstNames, booking.firstName)
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	//NOTE:You can't mix data types in Go Map
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("================================================")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
	fmt.Println("================================================")
	wg.Done() //NOTE:Decrease the waitgroup counter by 1
}
