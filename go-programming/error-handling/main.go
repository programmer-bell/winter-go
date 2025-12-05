package main

import "fmt"

func divide(a,b float64) (float64 , error) {
	if b == 0 {
		return  0 , fmt.Errorf("demonitor is not divided by zero")
	}
	return a/b, nil
}

func main(){
	fmt.Println("Error handling in go")
	ans, _ := divide(10,0)
	fmt.Printf("Division of the two number is %.2f",ans)
}

