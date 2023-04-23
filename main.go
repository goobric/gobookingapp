package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferentTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// placeholders in Go pkg.go.dev/fmt documentation
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and %v remaining to purchase.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name and tickets
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) //& pointer used to get the address of the variable

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to purchase")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets! You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
