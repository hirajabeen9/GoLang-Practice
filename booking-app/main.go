// Package main is the entry point of a Go program.
package main

// Importing the fmt package for formatted I/O operations.
import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//var bookings [50]string //array
	var bookings []string //slice

	//fmt.Printf("conferenceName: %T conferenceTickets: %T  remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Welcome to", conferenceName, "the booking app!")
	fmt.Printf("Get your %v tickets now! Only %v tickets are available!\n", conferenceName, conferenceTickets)
	fmt.Println(remainingTickets, "Remaining tickets!")

	var userName string
	var userBuyTickets uint
	var email string

	fmt.Println("Enter your name:")
	fmt.Scanln(&userName)
	fmt.Println("Enter your active email:")
	fmt.Scanln(&email)
	fmt.Println("Enter the number of tickets you want to buy:")
	fmt.Scanln(&userBuyTickets)

	remainingTickets = remainingTickets - userBuyTickets

	//bookings[0] = userName + " " + email
	bookings = append(bookings, userName+" "+email)
	fmt.Printf("Thank you %v for booking %v tickets for %v. you will receive confirmation details at this email %v \n", userName, userBuyTickets, conferenceName, email)
	fmt.Println("Be Quick!", remainingTickets, "tickets are remaining !")
	fmt.Printf("The Whole Array %v \n", bookings)
	fmt.Printf("The First Array %v \n", bookings[0])
	fmt.Printf("The Array Type %T \n", bookings)
	fmt.Printf("The Array Length/Size %v \n", len(bookings))

}
