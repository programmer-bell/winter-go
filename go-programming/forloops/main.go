package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Printf("The value is %d\n", i)
    }

	fruitsBuckets := [] string {"Apple", "Banana","Pinapple","Mango","Strabery"}

	for index, fruit := range fruitsBuckets{
		fmt.Printf("The pair is {%d : %s}\n",index,fruit)
	}


}