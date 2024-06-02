package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.facebook.com"

func main() {
	fmt.Println("Handling web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T", response)
	fmt.Println()
	// fmt.Println(response.Body)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
