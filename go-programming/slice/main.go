package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Number:",numbers)
	fmt.Printf("Number length is %d\n",len(numbers))
	numbers = append(numbers, 7,8,9,10,11,12,13,14,15)
	fmt.Println("Number:",numbers)
	fmt.Printf("Number has data type: %T\n",numbers)
	fmt.Printf("Number length is %d",len(numbers))

	stock := make([] int,0)
	fmt.Println("\nThe Slice:",stock)
	fmt.Println("Length:",len(stock))
	fmt.Println("Capacity:",cap(stock))

}
