package main

import "fmt"

func main() {
	fmt.Println("belajar switch")

	name := "putra"

	switch name {
	case "anggita":
		fmt.Println("hello ", name)
	case "putra":
		fmt.Println("hello ", name)
	default:
		fmt.Println("hello stranger")
	}

	switch length := len(name); length > 6 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("panjang nama sesuai")

	}

}
