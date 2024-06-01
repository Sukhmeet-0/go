package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	count := 23
	if count > 18 {
		fmt.Println("Adult")
	} else if count <= 18 {
		fmt.Println("Teen")
	} else {
		fmt.Println("Not allowed")
	}
}
