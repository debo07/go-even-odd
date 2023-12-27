package main

import "fmt"

func main() {
	numbers := make([]int, 10)

	// First create the slice by iterating
	for i := range(numbers) {
		numbers[i] = i + 1
	}

	for _, number := range(numbers) {
		if number % 2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}