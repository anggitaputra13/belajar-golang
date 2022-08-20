package main

import "fmt"

func main() {
	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	total := sumAll(1, 2, 3, 4, 5)
	total2 := kaliAll(numbers...)
	fmt.Println(total)
	fmt.Println(total2)

}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func kaliAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}
