package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	fmt.Println(add(5, 6))
	fmt.Println(adder(6, 7, 45, 3, 23, 2, 2, 34, 45, 6))
}

func add(a int, b int) int {
	return a + b
}

func adder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
