package main

import "fmt"


func simpleFunction(){
	fmt.Println("Greeting from simple function")
}


func add(a int,b int)  (result int) {
	result = a+b
	return
}


func main(){
	fmt.Println("Lets learn function")
	simpleFunction()
	ans := add(12,82)
	fmt.Printf("The result is %d",ans)
}



