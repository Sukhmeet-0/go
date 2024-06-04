package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("TWO")
	defer fmt.Println("THREE")
	defer fmt.Println("Hello")
	fmt.Println("World")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
