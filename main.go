package main

import "fmt"

func main() {

	var conferenceName = "Go Confirence"
	const conferenceTickets = 50
	var reminingTickets = 50

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
