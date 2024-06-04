package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Creating frontend server")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()

}
func PerformGetRequest() {
	const myurl = "http://localhost:1699/"
	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	// fmt.Println(response)
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func PerformPostRequest() {
	const myurl = "http://localhost:1699/post"
	requestBody := strings.NewReader((`
	{
		"name": "Sukhmeet",
        "age": 23,
        "gender": "male"
	}
	`))
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	// fmt.Println(response)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:1699/postform"
	data := url.Values{}
	data.Add("firstname", "sukhmeet")
	data.Add("lastname", "singh")
	data.Add("age", "22")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// fmt.Println(response)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
