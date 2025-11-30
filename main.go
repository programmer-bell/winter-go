package main

import "fmt"

func main() {

	// var conferenceName = "Go Confirence"
	// const conferenceTickets = 50
	// var reminingTickets = 50

	var conferenceName string = "Go Confirence"
	const conferenceTickets uint = 50
	var reminingTickets uint = 50

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

	// var userName = "Bell"
	// fmt.Printf("The user is :%v",userName)

	// var userName string
	// fmt.Scan(&userName)
	// fmt.Printf("The user is :%v",userName)

	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("Name:", name)

	// var name string
	// var age int
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Println("The username is:",name,"and age is:", age)

	var fname string
	var lname string
	var age uint
	var email string
	var userTickets uint

	fmt.Print("Enter your first name:")
	fmt.Scanln(&fname)
	fmt.Print("Enter your last name:")
	fmt.Scanln(&lname)
	fmt.Print("Enter your age:")
	fmt.Scanln(&age)
	fmt.Print("Enter your email address:")
	fmt.Scanln(&email)
	fmt.Print("Enter your booking tickets:")
	fmt.Scanln(&userTickets)

	reminingTickets -=  userTickets

	fmt.Printf("Thank you, %s %s. You are %d years old, your email is %s, and you have booked %d tickets.\n",
    fname, lname, age, email, userTickets)

	fmt.Printf("%v tickets remaining for %v\n",reminingTickets, conferenceName)


	}
