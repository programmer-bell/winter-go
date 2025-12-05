package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
	Age int
}


func main(){
	var prince Person
	fmt.Print("Prince person:",prince)
	prince.FirstName = "Prince"
	prince.LastName = "Albert"
	prince.Age = 21
	fmt.Print("\nPrince person:",prince)
	fmt.Printf("\nMr. %s %s is %d years old.",prince.FirstName,prince.LastName,prince.Age)

	person1 := Person{
		FirstName: "Akash",
		LastName: "Sutradhar",
		Age: 28,
	}
	fmt.Println("\nPerson 1 is :",person1)
}


