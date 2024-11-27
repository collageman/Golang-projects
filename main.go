package main

import (
	"fmt"
	"strings"
)

func main()  {
	conferenceName := "Phanerosis conference"
	const conferenceTickets = 50
	var remainingTickets uint=50

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	bookings :=[]string{}
	
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets= remainingTickets - userTickets
			bookings= append(bookings,firstName + " " + lastName)
			
			fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v shortly.\n", firstName, lastName, userTickets,email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			printFirstNames(bookings)
			
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("email addresss you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("number is tickets you entered is invalid")
			}
		} 
	}
}

func greetUsers(confName string, confTickets int, remainingTicket uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend.")

}

func printFirstNames (bookings []string) {
	firstNames := []string{}
			for _, booking := range bookings {
				names:=strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings: %v\n", firstNames)

}
/* 1. Add functionality to tell if the input is invalid after hitting enter and allow the user to enter again. 
2. the program to exit after 3 attempts. */
