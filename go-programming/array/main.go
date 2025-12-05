package main

import "fmt"

func main(){
	fmt.Println("We are learning array in golang")
	var name[5] string
	name[0] = "Python"
	name[1] = "JavaScript"
	name[2] = "Go"
	name[3] = "C++"

	fmt.Println("Names of the Person is :",name)

	var numbers = [5] int {1,2,3,4,5}
	fmt.Println("Number is :",numbers)

	fmt.Println("Length of Numbers array is :",len(numbers))

	fmt.Println("Value of 2nd index is:",name[1])

	var price [5] string
	fmt.Printf("Price is : %q",price)
}

