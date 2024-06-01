package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}

	roughvalue := 1
	for roughvalue < 10 {
		if roughvalue == 2 {
			goto lco
		}

		if roughvalue == 5 {
			roughvalue++
			continue
		}
		if roughvalue == 7 {
			break
		}

		fmt.Println(roughvalue)
		roughvalue++
	}
lco:
	fmt.Println(("Jumping"))
}
