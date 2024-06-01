package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you got 1")
		// fallthrough
	case 2:
		fmt.Println("You go 2")
		// fallthrough
	case 3:
		fmt.Println("You go 3")
		// fallthrough
	case 4:
		fmt.Println("You can go 4")
		// fallthrough
	case 5:
		fmt.Println("You can go 5")
		// fallthrough
	case 6:
		fmt.Println("Winner")
		// fallthrough
	default:
		fmt.Println("Invalid input")
	}
}
