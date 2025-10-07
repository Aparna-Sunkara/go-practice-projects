package main

import (
	"fmt"
	"practice/cal"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("Aparna")
	c, ok := cal.Add(5, 7)
	if !ok {
		fmt.Printf("I added and my values is : %d \n", c)
	} else {
		fmt.Printf("adding failed \n")
	}

	x := `10000
	
	3000`

	a := 100
	fmt.Println(string(a), x)

}
func init() {
	fmt.Println("Hello")

}

/*
func name(param1...) (response..){}
func - key word
name - name of the function
(param1...) - input params
(response..) - return values
*/
