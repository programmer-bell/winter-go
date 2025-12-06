package main

import (
	"fmt"
	"net/url"
)

func main(){
	fmt.Println("Learning URL")
	my_url := "https://jsonplaceholder.typicode.com/todos/?key=1&value=jon"
	fmt.Printf("Type of URL: %T\n",my_url)

	paredUrl , err := url.Parse(my_url)
	if err != nil{
		fmt.Println("Can't not parse URL",err)
	}
	fmt.Printf("Type of the URL: %T\n",paredUrl)
	fmt.Println("Schema:",paredUrl.Scheme)
	fmt.Println("Host:",paredUrl.Host)
	fmt.Println("Path:",paredUrl.Path)
	fmt.Println("RawQuery:",paredUrl.RawQuery)
}