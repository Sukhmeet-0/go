package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.September, 1, 10, 10, 10, 10, time.Local)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
