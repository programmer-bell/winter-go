package main

import "fmt"

func main(){
	x := 1
	if x > 5{
		fmt.Println("x is greater than 5")
	} else{
		fmt.Println("x is less than 5")
	}

	z := 210
	if z< 10{
		fmt.Println("The value is less than 10")
	} else if  z> 20{
		fmt.Println("The value is greater than 20")
	} else{
		fmt.Println("The value is was not found any idea")
	}
}



