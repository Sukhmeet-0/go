package main

import "fmt"

var name string = "Singh"

const LoginTOken string = "sdfklhjklas"

func main() {
	fmt.Println("Hello, 世界")
	var username string = "Sukhmeet singh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLogged bool = true
	fmt.Printf("Variable is of Type: %T \n", isLogged)

	var small uint8 = 255
	fmt.Println(small)

	var smallFloat float32 = 2.555555555555555555
	fmt.Println(smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)

	var website = "Sukhmeet singh"
	fmt.Println(website)

	numberOfUsers := 5
	fmt.Println(numberOfUsers, LoginTOken)

	fmt.Println(LoginTOken)
}
