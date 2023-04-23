package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferentTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// placeholders in Go pkg.go.dev/fmt documentation
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and %v remaining to purchase.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name and tickets
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //& pointer used to get the address of the variable

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you want to purchase: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thank you %v %v for booking %v tickets! You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are available for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Names of people who booked are: %v\n", firstNames)

		//check if remaining tickets is 0
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is fully booked! Please come back next year.")
			break
		}
	}

}
