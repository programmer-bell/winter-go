package main

import "fmt"

func main() {

	// var conferenceName = "Go Confirence"
	// const conferenceTickets = 50
	// var reminingTickets = 50

	var conferenceName string = "Go Confirence"
	const conferenceTickets int = 50
	var reminingTickets int = 50

	// fmt.Printf("ConferenceName is %T, conferenceTickets is %T and reminingTickets is %T \n" ,conferenceName,conferenceTickets,reminingTickets)

	// fmt.Print("Hello World\n")
	// fmt.Println("Hello World")
	// fmt.Print("Understand golang")

	// fmt.Println("Welcome to our ",conferenceName ,"booking application")

	fmt.Printf("Welcome to our %v booking application\n",conferenceName)

	// fmt.Println("We have total of",conferenceTickets, "tickets and",reminingTickets,"tickets are available")

	fmt.Printf("We have total of %v tickets and %v tickets are available\n",conferenceTickets,reminingTickets)
	fmt.Println("Get your tickets here to attend")

	// var name = "Go Confirence"
	// fmt.Print(name)

	}
