package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{} // slice

	fmt.Println("Welcome to", conferenceName, "the booking app!")
	fmt.Printf("Get your %v tickets now! Only %v tickets are available!\n", conferenceName, conferenceTickets)
	fmt.Println(remainingTickets, "Remaining tickets!")

	for {
		var userName string
		var userBuyTickets uint
		var email string

		fmt.Println("Enter your name (or type 'exit' to quit):")
		fmt.Scanln(&userName)

		// Exit condition if user types "exit"
		if strings.ToLower(userName) == "exit" {
			fmt.Println("Exiting the booking system. Thank you!")
			break
		}

		fmt.Println("Enter your active email:")
		fmt.Scanln(&email)
		fmt.Println("Enter the number of tickets you want to buy:")
		fmt.Scanln(&userBuyTickets)

		// Check if enough tickets are available
		if userBuyTickets > remainingTickets {
			fmt.Printf("Sorry, only %v tickets are remaining. Try again!\n", remainingTickets)
			continue // Skip the rest of the loop and ask for input again
		}

		remainingTickets -= userBuyTickets
		bookings = append(bookings, userName+" "+email)

		fmt.Printf("Thank you %v for booking %v tickets for %v. Confirmation sent to %v \n", userName, userBuyTickets, conferenceName, email)
		fmt.Println(remainingTickets, "tickets are remaining!")

		fmt.Printf("These are all our bookings: %v \n", bookings)
		fmt.Printf("***************************************************** \n")

		// Extract and print only customer names
		Names := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			Names = append(Names, names[0])
		}

		fmt.Printf("List of customer Names: %v \n", Names)

		// Stop booking if all tickets are sold
		if remainingTickets == 0 {
			fmt.Println("Sold out! No more tickets available.")
			break
		}
	}

	fmt.Println("Program ended. Thanks for using our booking system!")
}
