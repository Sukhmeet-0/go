package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello, 世界")
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you name: ")
	input, err := reader.ReadString('\n')
	fmt.Printf("You entered: %s", input)
	fmt.Print(err)

}
