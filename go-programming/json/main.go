package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `"json:name"`
	Age int `"json:age"`
	IsAdult bool `"json:is_adult"`
}

func main(){
	fmt.Println("We are learning json")
	person := Person{Name: "John",Age: 34,IsAdult: true}
	fmt.Println("Person data is :",person)

	// convert json encode (Marsel)
	jsonData , err := json.Marshal(person)
	if err != nil{
		fmt.Println("Error for joson conversion")
	}
	fmt.Println("Person data is :",string(jsonData))

	// convert json decode (Unmarsel)
	var personData Person
	err = json.Unmarshal(jsonData,&personData)
	if err != nil{
		fmt.Println("Error while unmarsel", err)
		return
	}
	fmt.Println("Person data is :",personData)
}
