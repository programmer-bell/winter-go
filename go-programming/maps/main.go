package main

import "fmt"


func main(){
	strudentGrades := make(map[string] int)

	strudentGrades["Prince"] = 31
	strudentGrades["Alice"] = 34
	strudentGrades["Bob"] = 36
	strudentGrades["Charlie"] = 38
	strudentGrades["Argan"] = 37

	fmt.Println("Marks of all_user:",strudentGrades)
	fmt.Println("Marks specific user:",strudentGrades["Bob"])
	
	delete(strudentGrades,"Bob")
	
	fmt.Println("Marks of all_user:",strudentGrades)
	fmt.Println("Marks specific user:",strudentGrades["Bob"])

	maggo, grade := strudentGrades["Alice"]
	fmt.Println("The user grade is",maggo)
	fmt.Println("The user exist:",grade)


	person := map[string] int {
		"Alice" : 90,
		"Bob" : 85,
		"Charlie" : 95,
	}

	for index, value := range person{
		fmt.Printf("Key is %s and marks is %d\n",index,value)
	}

}






