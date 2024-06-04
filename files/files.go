package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello, 世界")

	content := "This needs to go in file - sukhmeet.in"

	file, err := os.Create("./demo.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(file)
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()

	readFile("./demo.txt")
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}
