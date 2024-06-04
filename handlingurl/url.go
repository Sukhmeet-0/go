package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const murl = "https://www.google.com"

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("welcome to urls in golang")

	response, err := http.Get(murl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Body)
	defer response.Body.Close()

	result, _ := url.Parse(murl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Println(qparam)

	for _, val := range qparam {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
