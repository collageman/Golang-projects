package main

import (
    "booking-application/helper"
    "fmt"
)

const conferenceTickets = 50
const maxAttempts = 3

var conferenceName = "Phanerosis conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
    firstName       string
    lastName        string
    email           string
    numberOfTickets uint
}

func main() {
    greetUsers()
    
    attempts := 0
    for {
        if attempts >= maxAttempts {
            fmt.Println("Maximum number of invalid attempts reached. Program terminated.")
            break
        }

        firstName, lastName, email, userTickets := getUserInput()
        isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

        if isValidEmail && isValidName && isValidTicketNumber {
            bookTicket(userTickets, firstName, lastName, email)
            go helper.SendTicket(userTickets, firstName, lastName, email)

            firstNames := getFirstNames()
            fmt.Printf("The first names of bookings: %v\n", firstNames)

            if remainingTickets == 0 {
                fmt.Println("Our conference is booked out. Come back next year.")
                break
            }
            // Reset attempts counter after successful booking
            attempts = 0
        } else {
            attempts++
            remainingAttempts := maxAttempts - attempts
            
            if !isValidName {
                fmt.Println("First name or last name entered is too short.")
            }
            if !isValidEmail {
                fmt.Println("Email address you entered doesn't contain @ sign.")
            }
            if !isValidTicketNumber {
                fmt.Println("Number of tickets you entered is invalid")
            }
            
            if remainingAttempts > 0 {
                fmt.Printf("You have %d attempts remaining.\n", remainingAttempts)
            }
        }
    }
}

func greetUsers() {
    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
    firstNames := []string{}
    for _, booking := range bookings {
        firstNames = append(firstNames, booking.firstName)
    }
    return firstNames
}

func getUserInput() (string, string, string, uint) {
    var firstName string
    var lastName string
    var email string
    var userTickets uint
    
    fmt.Println("Enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Println("Enter your email address:")
    fmt.Scan(&email)

    fmt.Println("Enter number of tickets: ")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
    remainingTickets = remainingTickets - userTickets

    var userData = UserData{
        firstName:       firstName,
        lastName:        lastName,
        email:           email,
        numberOfTickets: userTickets,
    }

    bookings = append(bookings, userData)
    fmt.Printf("List of bookings is %v\n", bookings)

    fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v shortly.\n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}