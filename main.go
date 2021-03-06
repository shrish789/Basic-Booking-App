package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
// var bookings = make([]map[string]string, 0) // this is still a size. The size can automatically increazse or decrease depeending on use-case
var bookings = make([]UserData, 0)

func main() {

	// Greeting users
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName || !isValidEmail || !isValidTicketNumber || remainingTickets < userTickets {
			if !isValidName {
				fmt.Printf("First name or last name entered is not valid\n")				
			}
			if !isValidEmail {
				fmt.Printf("Email entered is not valid\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Ticket number entered is not valid\n")
			}
			if remainingTickets < userTickets {
				fmt.Printf("Total remaining tickets are %v and hence you can't book %v tickets\n", remainingTickets, userTickets)
			}
			continue
		}

		bookTickets(userTickets, firstName, lastName, email)
		wg.Add(1) // adds a thread and tells main thread that 1 more thread is running right now
		go sendTicket(userTickets, firstName, lastName, email)
		// writing go infront of a function forces it to run in a different thread.
		// Actually it is not thread rather its a goroutine which is better than normal kernel threads 
		// and hence concurrency management is better in GO as compared to other languages

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		
		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year\n")
			break
		}
		// fmt.Printf("These are all our bookings: %v\n", bookings)
	}
	wg.Wait() // makes the main thread to wait until all other threads are completed
}

func greetUsers() {
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
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
	// ask user for their name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// create a map for a user
	// var userData = make(map[string]string) // make is used to create a empty map
	
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10) // 10 is for decimal type
	
	bookings = append(bookings, userData)
	fmt.Printf("Bookings list: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(15 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done() // tells the main thread that one of the side threads is completed
}