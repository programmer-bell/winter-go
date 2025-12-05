package main

import (
	"fmt"
	"learngo/myutils"
)

func main(){
	fmt.Println("We are learning golang")
	myutils.PrintMessage("Hello World")

	var name string = "Badhan Sarkar"
	const age int = 21
	address := "Balarampur,Coochbehar"
	var salary float32 = 1200.502
	domain := "Backend Developer"


	fmt.Printf("A %v lives in %v,name Mr.%v who is %v years old, earned $%.2f in years",domain,address,name,age,salary)




}
