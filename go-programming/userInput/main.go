package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Print("What is your name:")
	// var name string
	// fmt.Scan(&name)
	// fmt.Printf("Hello Mr. %s",name)
	
	
	fmt.Print("What is your name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello Mr. %s",name)
}
