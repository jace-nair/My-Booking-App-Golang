package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceTickets int = 100
	var remainingTickets uint = 100
	var conferenceName string = "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// asking for user input
		fmt.Println("Please enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Please enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Please enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Please Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		// validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			// book ticket in system
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// print only first names
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names %v\n", firstNames)

			// exit application if no tickets are left
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}

}
