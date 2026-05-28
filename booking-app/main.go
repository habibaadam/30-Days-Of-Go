package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50


	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v conferenceTickets and a remaining of %v tickets are still available\n", conferenceTickets, remainingTickets)
	//fmt.Printf("conference name is %T, conference tickets is %T and remaining tickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	bookings := []string{}  // dynamic slice without a fixed size


	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

	    //var bookings [50]string  // arrays defined with data type and fixed size


		fmt.Println("Please Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please Enter your first email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you would want: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

	    //bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)
		fmt.Printf("Thank you %v %v for booking %v number of tickets. You will receive an email at %v to confirm details of your booking.\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

        firstNames := []string{}
		for _, booking := range(bookings) {
			var names = strings.Fields(booking) // used to seperate words based on whitespace
			firstNames = append(firstNames, names[0])
		}

		

		fmt.Printf("All the bookings made are: %v\n", firstNames)
	}

}