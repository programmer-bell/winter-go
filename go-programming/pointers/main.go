package main

import "fmt"

func main(){
	// var num int
	var num int = 2

	// var ptr *int
	// ptr = &num

	ptr := &num

	fmt.Println("Num has value:",num)
	fmt.Println("\nPointer contains:",ptr)
	fmt.Println("\nPointer contains:",*ptr)
	
	*ptr = 12
	lol := *ptr
	fmt.Println("The new value is:",lol)
	fmt.Println("The num is:",num)
	fmt.Println("\nPointer contains:",ptr)


}
