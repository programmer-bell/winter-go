package main

import "fmt"

func main() {
	var num int = 42
	fmt.Println("Number is",num)
	fmt.Printf("Type of num is %T\n",num)

	data := float64(num)
	fmt.Printf("Data is %.2f\n",data)
	fmt.Printf("Type of data is %T\n",data)
}

