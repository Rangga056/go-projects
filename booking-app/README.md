# Go Conference Booking Application

This is a simple conference ticket booking application written in Go. It allows users to book tickets for a conference, validating their input and sending a confirmation email asynchronously using goroutines.

## Features

- User input validation for name, email, and number of tickets.
- Asynchronous ticket confirmation email sending using goroutines.
- Concurrent handling of multiple bookings with `sync.WaitGroup`.
- Simple CLI for booking tickets and viewing current bookings.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installing

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-conference-booking.git
   cd go-conference-booking

   ```

2. Run the application:

```bash
    go run main.go
```

### Usage

1. Run the application

```bash
    go run main.go
```

2. Follow the prompts to enter your first name, last name, email, and number of tickets you wish to book.

3. If the input is valid, your booking will be confirmed, and a confirmation message will be displayed. The application will also simulate sending a ticket to your email address asynchronously.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Go](https://golang.org/) - The Go programming language
- [Go Redis](https://github.com/go-redis/redis) - Redis client for Go
- Inspiration from various Go tutorials and examples.
