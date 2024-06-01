package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// no inheritance in golang: no super or parent

	sukhmeet := User{"Sukhmeet singh", "xsukhmeet@gmail.com", true, 22}

	fmt.Printf("Details are: %+v\n", sukhmeet)
	fmt.Printf("Name is %v", sukhmeet.Name)
}

type User struct {
	Name   string
	Email  string
	Statut bool
	Age    int
}
