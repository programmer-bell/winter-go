package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file , error := os.Create("example.txt")
	if error != nil{
		fmt.Println("Error while creating file:",error)
		return 
	}
	defer file.Close()

	content := "Hello world in new file by go"
	_, error = io.WriteString(file,content)
	if error != nil{
		fmt.Println("Error while wrighting a foile")
	}

	fmt.Println("successfully created file")
}
