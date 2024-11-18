package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"

var remainingTickets uint = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // Go routine: by using go keyword we are creating go routines which are go threads (Green Thread) and not one complete OS thread which is very expensive. with one OS thread we can create multiple Green Threads. Green threds are managed by Go runtime. In languages like Java 1 OS thread = 1 Java thread. which is very expensive.
// Go routines also have a concept of "Channels" which facilitates the communication between Green Threads which is not easily achievable in other Programming langauges

//var bookings = make([]map[string]string, 0) // this is a map declaration and initialization [0 in the make func is the initialization parameter and since slice grows dynamically we can provide initial size as 0]
// var bookings [50]string  : this is a array
// var bookings []string // this is a slice initialized with zero values in it : var bookings []string{}

func main() {
	// fmt.Print("Hello World")

	//var bookings = []string{} // alternate slice declaration 1
	// bookings := []string{} // alternate slice declaration 2

	// fmt.Println("This prints value of remainingTickets ", remainingTickets)
	// fmt.Println("This prints reference of &remainingTickets ", &remainingTickets)
	// fmt.Println("Welcome to", conferenceName, "booking applicaton")

	greetUsers()

	// fmt.Printf("conferenceTicket is %T, remianingTicket is %T, conference is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// var bookings = [50]string{"Nana", "vijay", "peter"}

	// for{ } [or this for true {}] , for condition {}, for i:=0, condition, i++{}

	// for remainingTickets > 0 && len(bookings) < 50 {

	// Get user Input function call
	firstName, lastName, email, userTickets := getUserInput()

	// UserValidation Function Call

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// isValidCity := city =="Singapore" || city == "London"

	// if userTickets > remainingTickets {
	// 	fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets", remainingTickets, userTickets)
	// 	continue
	// }
	if isValidName && isValidEmail && isValidTicketNumber {
		//  book ticket function call

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)                                              // 1 is number of threads to wait for. If we have another go func() we would have entered 2 or for how many threads we want the main thread to wait for until their completetion donot exit the main thread.
		go sendTicket(userTickets, firstName, lastName, email) // here go keyword start a new thread which is starts a new goroutine

		// fmt.Printf("The whole slice %v \n", bookings)
		// fmt.Printf("The first value %v \n", bookings[0])
		// fmt.Printf("The first type %T \n", bookings)
		// fmt.Printf("The Slice size is %v \n", len(bookings))

		// Print first Names Function

		firstNames := getFirstNames()
		fmt.Printf("The firstNames of bookings are: %v \n", firstNames)

		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {

			// end program

			fmt.Print("Our conference is booked out. Come back next year.")
			// break  // use break when you enable for loop. we disabled for loop to learn about weighed groups which helps with blocking main thread until Green Threads are completed execution.
		}
	} else {
		// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

		if !isValidName {
			fmt.Println("first name or last name you enterd is too short, Please try again.")
		}
		if !isValidName {
			fmt.Println("Email address you entered does not contail @ sign, Please try again.")
		}
		if !isValidName {
			fmt.Println(" Number of tickets you entered is invalid, Please try again.")
		}

	}
	wg.Wait()

	// }

	// city := "London"

	// switch city {

	// case "New york":
	// 	// code
	// case "Singapore","Mexico City":
	// 	// code
	// default:
	// 	fmt.Println("No valid city selected")
	// }

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	//for-each loop
	// _ are use to explicityly say we are not using that variable which nor mally requires like for index , booking := range bookings

	for _, booking := range bookings {

		// var names = strings.Fields(booking)
		// var firstName = names[0]

		//firstName := booking["firstName"] //  var firstName string= booking["firstName"]
		firstName := booking.firstName

		firstNames = append(firstNames, firstName)

	}
	// fmt.Printf("The whole slice %v \n", bookings)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask uer for their name

	// userName = "Tom"
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter your number of tickets to be booked : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// declaration of slice looks like :  var myslice []string
	// declaration of map looks like :  var mymap [string]string
	// Create a map for a user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	//var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation on you email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets reamining out of %v for the %v\n", remainingTickets, conferenceTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("##################")

	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)

	fmt.Println("##################")

	wg.Done() // done function removes the thread fromt he waiting list which is checked by wg.Wait(). Once the all the threads in the wg.Wait() is completed, then the main thread will be terminated.
}
