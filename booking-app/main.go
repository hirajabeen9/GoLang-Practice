// Package main is the entry point of a Go program.
package main

// Importing the fmt package for formatted I/O operations.
import "fmt"

func main() {
	// Declaring a string variable using the var keyword.
	var conferenceName = "Go Conference"

	// Declaring a constant using the const keyword.
	// Constants are immutable and cannot be changed after declaration.
	const conferenceTickets = 50

	// Declaring a variable using the var keyword.
	// Variables declared with var can be reassigned later.
	var remainingTickets = conferenceTickets - 10

	// Printing a welcome message using fmt.Println, which adds a newline at the end.
	fmt.Println("Welcome to", conferenceName, "the booking app!")

	// Printing a formatted string using fmt.Printf.
	// %v is a placeholder that gets replaced with the corresponding variable value.
	fmt.Printf("Get your %v tickets now! Only %v tickets are available!\n", conferenceName, conferenceTickets)

	// Printing the number of remaining tickets.
	fmt.Println(remainingTickets, "Remaining tickets!")

	var userName string
	var userBuyTickets int
	fmt.Println("Enter your name:")
	fmt.Scanln(&userName)
	fmt.Println("Enter the number of tickets you want to buy:")
	fmt.Scanln(&userBuyTickets)
	fmt.Printf("Thank you %v for booking %v tickets for %v!\n", userName, userBuyTickets, conferenceName)

}
