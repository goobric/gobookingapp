package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferentTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// placeholders in Go pkg.go.dev/fmt documentation
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("There are %v tickets in total, and %v remaining to purchase.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name and tickets

	userName = "Nana"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets!\n", userName, userTickets)

}
