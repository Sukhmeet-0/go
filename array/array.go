package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	var arr [4]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	fmt.Println(arr)
	fmt.Println(len(arr))

	var veglist = [3]string{
		"Tomato",
		"Potato",
		"Onion",
	}
	fmt.Println(veglist)
}
