package main

import (
	"fmt"
	"time"
)


func main(){
	// yy-mm-dd hh-mm-ss

	currentTime := time.Now()
	fmt.Println("Current time:",currentTime)
	fmt.Printf("Type of currentTime %T\n",currentTime)

	format := currentTime.Format("02-01-2006, Monday,15:04:05")
	fmt.Println("Formatted time:",format)

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time,_ := time.Parse(layout_str,dateStr)
	fmt.Println("Formetted time:",formatted_time)

}

