package main

import (
	"fmt"
	"strings"
)

func main() {

	// var conferenceName = "Go Confirence"
	// const conferenceTickets = 50
	// var reminingTickets = 50

	var conferenceName string = "Go Confirence"
	const conferenceTickets uint = 50
	var reminingTickets uint = 50

	// var booking = [50] string {"Jone","Nana","Alice"}
	var booking [] string


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

	for{

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

			// var isValidName bool = len(fname) >= 2 && len(lname) >= 2
			isValidName := len(fname) >= 2 && len(lname) >= 2
			isValidEmail := strings.Contains(email,"@")
			isValidTicketNumber := userTickets > 0  && userTickets <= reminingTickets

			if isValidName && isValidEmail && isValidTicketNumber {

					reminingTickets -=  userTickets

					// booking[0] = fname + " " +lname
					booking = append(booking,fname + " " +lname )

					// fmt.Printf("The whole array is :%v\n",booking)
					// fmt.Printf("The first array is: %v\n",booking[0])
					// fmt.Printf("The array type is: %T\n",booking[0])
					// fmt.Printf("The array length is: %d\n",len(booking))

					fmt.Printf("The whole slice is :%v\n",booking)
					fmt.Printf("The first slice is: %v\n",booking[0])
					fmt.Printf("The slice type is: %T\n",booking[0])
					fmt.Printf("The slice length is: %d\n",len(booking))

					fmt.Printf("Thank you, %s %s. You are %d years old, your email is %s, and you have booked %d tickets.\n",
					fname, lname, age, email, userTickets)

					fmt.Printf("%v tickets remaining for %v\n",reminingTickets, conferenceName)

					firstNames := []string{}
					for _, booking := range booking{
						var names = strings.Fields(booking)
						firstNames = append(firstNames,names[0])
					}

					fmt.Printf("These are all about our booking application:%v",booking)

					// var noTicketRemaining bool =  reminingTickets == 0 
					
					// noTicketRemaining := reminingTickets == 0

					// if noTicketRemaining {
					if reminingTickets == 0 {
						fmt.Println("Our conference is booked out. Come back next year.")
						break
					}

			} else{

				// fmt.Printf("We only have %v ticktes remaining .So,you can not but %v tickets.",reminingTickets, userTickets)
				// continue

				if !isValidName {
					fmt.Println("first name or last name is too short")
				}
				if !isValidEmail{
					fmt.Println("email address you entered doesnot contain @ sign")
				}
				if !isValidTicketNumber{
					fmt.Println("number of tickets you enter is invalid")
				}

			}


		}

}


