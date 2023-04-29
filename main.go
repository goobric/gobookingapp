package main

import (
	"fmt"
	"myGoBooking/common"
	"sync"
	"time"
)

const confTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = [50]string{} // an array of strings with max 50 users
// var bookings = make([]map[string]string, 0) // an empty list of Maps; must be initialized with a starting value
var bookings = make([]UserData, 0) // an empty list of Structs
// create the Struct with the Type keyword, name the Struct,  define the key and its Data Type
type UserData struct {
	firstName  string
	lastName   string
	email      string
	numOfTicks uint
}

var wg = sync.WaitGroup{}

func main() {
	// replaced the fmt.Print lines with a Function Call
	greetUsers() // calls the greetUsers function and prints 'Welcome All'

	// fmt.Printf("confTickets is %T, remainingTickets is %T, conferenceName is %T\n", confTickets, remainingTickets, conferenceName)

	// remove the for-loop  {}
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketAmount := common.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// isValidCity := city == "Singapore" || city == "Berlin"
	// !isValidCity

	if isValidName && isValidEmail && isValidTicketAmount {

		bookedTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicketInfo(userTickets, firstName, lastName, email)
		// Call function printFirstNames
		firstNames := getFirstNames()
		fmt.Printf("The people who have tickets are: %v\n", firstNames)

		// var noTicketsRemaining bool = remainingTickets == 0
		if remainingTickets == 0 {
			// end program
			fmt.Println("Conference is fully booked. Come back next time.\n")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("first or last name is too short")
		}
		if !isValidEmail {
			fmt.Println("email doesn't contain @ sign")
		}
		if !isValidTicketAmount {
			fmt.Println("ticket amount is invalid")
		}
		// fmt.Printf("There are only %d tickets remaining, you can't book %v tickets.\n", remainingTickets, userTickets)
		// fmt.Println("Your input data is invalid, please try again.\n")
	}
	wg.Wait()
}

// passing Conference Name, Conference Tickets and Remainging Tickets into this new function as Input Parameters
// When the function is Called, all of the parameters have to be passed.
func greetUsers() {
	fmt.Println("Welcome to the %v booking program\n", conferenceName)
	fmt.Printf("There are a total of %v tickets and %v tickets are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
}

func getFirstNames() []string {
	firstNames := []string{} // a slice of first names
	for _, booking := range bookings {

		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	// fmt.Printf("People who booked tickets: %v\n", firstNames)
	return firstNames
}

// validateUserInput function is common across all Go Conference Venues
// this function is stored in a file called common.go

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user to enter their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets required: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookedTicket(userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	// update remaining tickets available
	remainingTickets = remainingTickets - userTickets

	// create a map for the a user
	// var userData = make(map[string]string)
	//create a struct for the user
	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		numOfTicks: userTickets,
	}
	/* userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	// convert userTickets uint to a String data type, using Goland buildin functions.
	userData["numOfTicks"] = strconv.FormatUint(uint64(userTickets), 10) */
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("1st value of array: %v\n", bookings[0])
	// fmt.Printf("Slice data type: %T\n", bookings)
	// fmt.Printf("Length of slice: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirm email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// simulation of sending ticket information to a user via email.
func sendTicketInfo(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticketinfo = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("Sending ticket information:\n %v \nto email: %v\n", ticketinfo, email)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	wg.Done()
}
