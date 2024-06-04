package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// no inheritance in golang: no super or parent

	sukhmeet := User{"Sukhmeet singh", "xsukhmeet@gmail.com", true, 22, 18}

	fmt.Printf("Details are: %+v\n", sukhmeet)
	fmt.Printf("Name is %v", sukhmeet.Name)
	fmt.Println()
	sukhmeet.GetStatus()
	sukhmeet.email()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) email() {
	u.Email = "sukhmeets111@gmail.com"
	fmt.Println("Email is: ", u.Email)
}
