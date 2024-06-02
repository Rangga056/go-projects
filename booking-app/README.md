# Go Conference Booking Application

## Overview

This Go program simulates a simple conference ticket booking system. It allows users to book tickets, validates their input, and keeps track of the remaining tickets.

## Features

- Book tickets for a conference.
- Validate user input for names, email, and number of tickets.
- Track the remaining tickets.
- Display the first names of all users who have booked tickets.

## How It Works

### Package Declaration

The `main` package is the starting point of a Go program. The `main` function within this package is the entry point of the program.

### Imports

- `fmt`: Implements formatted I/O with functions analogous to C's printf and scanf.
- `strings`: Contains functions to manipulate UTF-8 encoded strings.

### Global Variables and Constants

- `conferenceTickets`: A constant representing the total number of tickets available.
- `conferenceName`: A variable holding the name of the conference.
- `remainingTickets`: A variable representing the number of tickets still available.
- `bookings`: A slice that stores the names of people who have booked tickets.

### Main Function

1. The `main` function starts by greeting the user.
2. It enters a loop that continues as long as there are remaining tickets and the number of bookings is less than 50.
3. The loop prompts the user for their booking details and validates the input.
4. If the input is valid, it books the ticket and displays the first names of all bookings.
5. If no tickets remain, it prints a message and breaks the loop.

## Running the Program

1. Install Go from [golang.org](https://golang.org/).
2. Copy the code into a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the program with the command:
   ```bash
   go run main.go
   ```

## Example Usage

```bash
Welcome to Go Conference booking application
We have a total of 50 tickets and 50 are still available.
Get your tickets here to attend

Enter your first name:
John
Enter your last name:
Doe
Enter your email:
john.doe@example.com
Enter your number of tickets:
2
Thank you John Doe for booking 2 tickets. You will receive a confirmation email at john.doe@example.com
48 tickets remaining for Go Conference

These are all our bookings: [John]
```

This `README.md` provides an overview of the Go conference booking application, explains how it works, and guides you on how to run the program.
