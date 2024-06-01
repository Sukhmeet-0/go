package main

import "fmt"

func main() {
	// fmt.Println("Hello sukh, welcome to pointers")
	// var ptr *int

	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println(myNumber)

}
