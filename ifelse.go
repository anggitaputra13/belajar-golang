package main

import "fmt"

func main() {
	name := "anggita"

	if name == "anggita" {
		fmt.Println("hello ", name)
	} else if name == "putra" {
		fmt.Println("hello ", name)
	} else {
		fmt.Println("hello stranger")
	}

	if length := len(name); length > 6 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("panjang nama sesuai")

	}
}
