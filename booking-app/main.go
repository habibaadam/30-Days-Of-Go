package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
//var bookings = make([]map[string]string, 0) // list of maps


var bookings = make([]UserInfo, 0)
type UserInfo struct {
	firstName string
	lastName string
	email string
	numOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	//var bookings [50]string  // arrays defined with data type and fixed size

	//for {
		firstName, lastName, email, userTickets := getUserInput()

		// basic user input validation
		isValidName, isValidEmail, isValidUserTicket := helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)
        if isValidName && isValidEmail && isValidUserTicket {
			bookings = bookTickets(userTickets, firstName,lastName, email)

			wg.Add(1) // specifies go thread to be waited for
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := extractFirstNames()

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked! Come back next year\n")
				//break // ends the indefinite loop
			}

			fmt.Printf("All the bookings made are: %v\n", firstNames)
		} else {
			if !isValidName {
				fmt.Println("Your first name and last name must be more than 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Incorrect email format detected")
			}
			if !isValidUserTicket {
				fmt.Println("Number of tickets requested is invalid")
			}
		}
		wg.Wait() // wait for the goroutine in counter to be 0
	}
//}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v conferenceTickets and a remaining of %v tickets are still available\n", conferenceTickets, remainingTickets)
}

func getUserInput() (string, string, string, uint) {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please Enter your first email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you would want: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func extractFirstNames() []string {
	firstNames := []string{}
	for _, booking := range(bookings) {
		//var names = strings.Fields(booking) // used to seperate words based on whitespace
		firstNames = append(firstNames, booking.firstName)
		}
	return firstNames
}

func bookTickets(userTickets uint ,firstName string, lastName string, email string ) []UserInfo {
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = firstName + " " + lastName

	var userInfo = UserInfo {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numOfTickets: userTickets,
	}


	bookings = append(bookings, userInfo)
	//fmt.Printf("List of all bookings are %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v number of tickets. You will receive an email at %v to confirm details of your booking.\n", firstName, lastName, userTickets, email)
				fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return bookings
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	// simulating delay for 15 seconds
	time.Sleep(15 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("----------------------------------")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("------------------------------------")
	wg.Done()
}