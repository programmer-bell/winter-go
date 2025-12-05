package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Apple,orange,banana"
	parts := strings.Split(data,",")
	fmt.Println(parts)

	str := "One or two or three or fibe"
	count := strings.Count(str,"or")
	fmt.Println("count:",count)

	str = "     Hello World    "
	fmt.Println(str)
	trim := strings.TrimSpace(str)
	fmt.Println(trim)

	str1 := "Badhan"
	str2 := "Sarkar"
	resutl := strings.Join([]string{str1,str2},"-")
	fmt.Println("result:",resutl)
}