package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	languages := make(map[string]string)

	languages["Js"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages ", languages)

	for key, value := range languages {
		// fmt.Println(key, value)
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
