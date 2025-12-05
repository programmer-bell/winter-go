package main

import (
	"fmt"
	"learngo/myutils"
)

func main(){
	fmt.Println("We are learning golang")
	myutils.PrintMessage("Hello World")

	var name string = "Mr. Allen"
	const age int = 00
	address := "Not Found"
	var salary float32 = 120.09
	domain := "Backend Developer"


	fmt.Printf("A %v lives in %v,name Mr.%v who is %v years old, earned $%.2f in years",domain,address,name,age,salary)




}
