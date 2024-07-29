package main

import "fmt"

func main() {
	divisibleByThree()
}

var number = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var num int

func divisibleByThree() {
	for index := 0; index < len(number); index++ {
		if number[index]%3 == 0 {
			num = number[index]
		}
		fmt.Println(num)
	}
}
