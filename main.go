package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Println("There are", conferenceTickets, "tickets in total, and ", remainingTickets, " remaining to purchase.")
	fmt.Println("Get your tickets here to attend")

}
