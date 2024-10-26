package main

import (
	"fmt"
	"strings"
)

func main() {

	//---------------START THE APP--------------------------
	//var nameOfVariable typeOfData = value
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//define a slice variable - name of variable & type of value it will hold along with memory allocation.
	var bookings []string
	//define a array variable - name of variable & type of value it will hold along with memory allocation.
	//var bookings []string

	for {

		// user name and tickets
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//-------------ASKING FOR USER INPUT (SCAN)------------------------
		// Get the user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		//array
		//bookings[0] = firstName + "" + lastName

		//slice
		bookings = append(bookings, firstName+""+lastName)

		//-------------PRINT THE SUMMARY OF BOOKINGS--------------------------
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		//-------------EXIT THE APP BY PRINTING SUMMARY MESSAGE FOR THE USER-----------------

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		//Define slice only for First Names
		firstNames := []string{}

		// Loop through or iterate through bookings list and grab firstName at a time
		for _, booking := range bookings { // Get the index of the element and the value of the element from the bookings slice. The value will have firstName and lastName with a space character in between.
			var names = strings.Fields(booking)       // Split booking value string on a space character using the Fields function from strings package. Results in a slice with two elements. Save it as names array.
			firstNames = append(firstNames, names[0]) //Get the value of element at the index 0 and append it into firstNames slice
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

	}

}
