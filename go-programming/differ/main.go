package main

import (
	"fmt"
)


func add(a,b int) int {
	return  a+b 
}


func main(){
	defer fmt.Println("Starting of the program")
	data := add(5,6)
	fmt.Println("Data is:",data)
	defer fmt.Println("Middle of the program")
	defer fmt.Println("End of the program")
}
