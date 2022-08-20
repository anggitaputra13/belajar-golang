package main

import (
	"fmt"
)

func main() {
	fmt.Println("perulangan")
	// for
	// for counter := 1; counter <= 20; counter++ {
	// 	if counter%3 == 0 && counter%5 == 0 {
	// 		fmt.Println("fizz & buzz")
	// 	} else if counter%3 == 0 {
	// 		fmt.Println("fizz")
	// 	} else if counter%5 == 0 {
	// 		fmt.Println("buzz")
	// 	} else {
	// 		fmt.Println(counter)

	// 	}
	// }

	// for range
	slice := []string{
		"aku",
		"kamu",
		"dia",
		"mereka",
	}

	person := make(map[string]string)
	person["name"] = "anggita"
	person["age"] = "22"

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }
	for index, value := range slice {
		if index == 0 {
			continue
		} else if index == len(slice)-1 {
			break
		} else {
			fmt.Println(index, "=", value)

		}
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
