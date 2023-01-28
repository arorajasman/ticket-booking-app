package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"ticket-booking-app/helper"
	"time"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}
var userBookings = make([]map[string]string, 0)

type UserDetails struct {
	firstName     string
	lastName      string
	email         string
	ticketsBooked uint
}

var listOfUsersBookingForConference = make([]UserDetails, 0)

// ! using the WaitGroup Structure from the sync packge to create a wait group to let the
// ! main thread to wait to complete its execution until all the threads / go routines inside
// ! the main method has completed their execution
var wg = sync.WaitGroup{}

func main() {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	greetUser()

	fmt.Println("Get your tickets here to attend")

	// ! using the infinite for loop to create a loop to get the details to book the tickets
	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&userFirstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&userLastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&userEmail)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, userFirstName, userLastName, userEmail)

			// ! ----------------------------------------------------------------------------------
			// ! calling the sendEmail() function just to simulate the concurrency using go routines in go
			data := UserDetails{firstName: userFirstName, lastName: userLastName, ticketsBooked: userTickets, email: userEmail}

			fmt.Println("starting the thread for sendEmail() function execution")

			// ! using the wg instance to get access to the Add() method to add the number
			// ! of threads / go routines for which we want the main method to wait to complete
			// ! their execution before the main methods completes its execution and stops the program

			// ! the add method should be called before calling any of the go routines and should
			// ! be provided with the number of go routines that we have in the file
			// ! since we only hace one goroutine for sending the email so we are adding only
			// ! 1 to the Add() method as input
			wg.Add(1)

			// ! adding go keyword in front of sendEmail to create a seperate thread
			// ! from the main thread when the program tries to execute the sendEmail()
			// ! function so that the main function / thread will not get blocked while
			// ! the sendEmail() function takes some time to execute

			// ! the go keyword below is used to start a new thread / go routine only for the
			// ! execution of the sendEmail() function everytime the flow of code hits the sendEmail()
			// ! function

			// ! till the time of execution of sendEmail() this sendEmail() thread will be working
			// ! in the back of the application

			go sendEmail(data)

			// ! ----------------------------------------------------------------------------------

			fmt.Printf("These are all our bookings: %v\n", bookings)

			var firstNames []string = getFirstNames()

			fmt.Println("The first names of bookings are:", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Conference is booked out. Come back next year")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("Either first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain the @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

		}

	}

	// ! using the Wait() method from the wg instance to wait for all (number of) the go routines
	// ! that we have added in the Add() method to complete their execution before the
	// ! main() method completes its execution

	// ! the Wait() method should always be called at the end before closing bracket of
	// ! the main method
	wg.Wait()

}

// ! we do not need to pass the conferenceName, conferenceTickets and remainingTickets as
// ! input since these are defined at the package level
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
}

// ! we do not need to pass the bookings slice as input since it is defined at the
// ! package level and is accessible to all the functions defined in the package
func getFirstNames() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		name := strings.Fields(booking)
		firstNames = append(firstNames, name[0])

	}
	return firstNames
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	fmt.Println("The data inside the userData map is: ")
	fmt.Println(userData)

	userBookings = append(userBookings, userData)

	fmt.Println("The data inside the userBookings slice is: ")
	fmt.Println(userBookings)

	details := UserDetails{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		ticketsBooked: userTickets,
	}

	fmt.Println("the details of user stored in the structure UserDetails is: ")
	fmt.Println(details)

	fmt.Println("First name of the user bookings the ticket is:", details.firstName)

	listOfUsersBookingForConference = append(listOfUsersBookingForConference, details)

	fmt.Println("the data inside the listOfUsersBookingForConference slice is:")
	fmt.Println(listOfUsersBookingForConference)

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will recieve a confirmation by email at", email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}

// ! the code below is just for example to show concurrency in go using goroutines
func sendEmail(userDetails UserDetails) {
	// ! the Sleep() method below stops the execution of the thread / goroutine for 10 seconds
	time.Sleep(10 * time.Second)

	fmt.Println("***************************************************************")
	fmt.Println(userDetails.ticketsBooked, "tickets for", userDetails.firstName, userDetails.lastName)
	fmt.Println("Sending ticket to", userDetails.email)
	fmt.Println("***************************************************************")

	// ! using the Done() method here to remove this sendEmail() thread and decrement
	// ! the value stored in Add() method when this method has done its execution as a
	// ! seperate thread from the main thread

	// ! the Done() method is added only inside method that we use a seperate goroutine / thread
	wg.Done()
}
