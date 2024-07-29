package main

import "fmt"

func main() {

	maxAndMin()
}

var maximum int
var minimum int
var numbers = [10]int{1, 3, 45, 67, 89, 7, 65, 32, 4, 10}

func maxAndMin() {
	for index := 0; index < len(numbers); index++ {
		if numbers[index] > maximum {
			maximum = numbers[index]
			fmt.Println(maximum)
		} else if numbers[index] < minimum {
			minimum = numbers[index]
			fmt.Println(minimum)
		}

	}
}
