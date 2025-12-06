package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web request..")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error getting GET response")
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response %T\n", res)
	fmt.Println("response:", res)

	data , error := io.ReadAll(res.Body)
	if error != nil{
		fmt.Println("Error reading response:",error)
		return
	}
	fmt.Println("response:",string(data))
}
