package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 23
	name := "anggita"
	// result := sayHello(age, name)
	data1, _ := sayHello2(age, name)
	fmt.Println(data1)
}

func sayHello(angka int, data string) string {
	result := "nama: " + data + ", umur: " + strconv.FormatInt(int64(angka), 10)
	return result
}

func sayHello2(angka int, data string) (string, int) {
	return data, angka
}
