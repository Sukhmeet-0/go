package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, 世界")

	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	// fruitlist = append(fruitlist, "Orange", "Banana", "Mango")
	fruitlist = append(fruitlist, "Orange", "Banana", "Mango")
	fmt.Println(fruitlist)

	fruitlist = fruitlist[1:3]
	fmt.Println(fruitlist)

	highscore := make([]int, 4)
	highscore[0] = 100
	highscore[1] = 200
	highscore[2] = 300
	highscore[3] = 400
	// highscore[4] = 500 // index out of bound
	highscore = append(highscore, 555, 666, 777)

	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	// remove  avalue from slices

	var courses = []string{"React", "Java", "Python", "Go", "Rust"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
