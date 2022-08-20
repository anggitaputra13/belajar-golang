package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Anggita",
		"age":  "23",
	}

	fmt.Println(person)
	fmt.Println(len(person))
	delete(person, "age")
	fmt.Println("person2", person)
	fmt.Println("person2", len(person))

}
