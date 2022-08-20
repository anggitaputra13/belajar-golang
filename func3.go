package main

import "fmt"

type Filter func(string) string

func main() {
	callKata("anggita", filterKata)

}

func callKata(name string, filter Filter) {
	result := filter(name)
	fmt.Println("Hallo", result)
}

func filterKata(kata string) string {
	if kata == "anjing" {
		kata = ""
	}

	return kata
}
