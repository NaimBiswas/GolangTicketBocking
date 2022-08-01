package main

import "fmt"

func main() {
	//var helloAnotherWorld = "Hello from another world"
	//var helloAnotherUniverse = "Hello from another universe"
	const totalTickets = 50
	var remainingTickets = 50
	// var nameConforencer = "Naim Biswas"
	var userName string
	var emailAbrress string
	var numberOfTicket int
	fmt.Printf("Hello World, Remaining: %v  total: %v \n", remainingTickets, totalTickets)

	fmt.Println("Enter Your Name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter Your Email: ")
	fmt.Scan(&emailAbrress)

	fmt.Println("Enter Ticket that you want: ")
	fmt.Scan(&numberOfTicket)

	fmt.Printf("Thank %v  for booking %v tickets. We will send you a confirmation mail on %v \n", userName, numberOfTicket, emailAbrress)
	remainingTickets = remainingTickets - numberOfTicket
	fmt.Println("Remaining tickets are: ", remainingTickets)
}
