package main

import "fmt"

func main() {
	salesCalculator()
}

var itemsSold int
var reply string
var percentage int
var commission int

func salesCalculator() {
	for index := 0; index < 5; index++ {
		fmt.Println("Enter the items you sold ")
		fmt.Scan(&itemsSold)

		fmt.Println("Enter the percentage ")
		fmt.Scan(&percentage)

		commission = itemsSold * percentage / 100
		fmt.Println("Do you have other items you've sold? ")
		fmt.Scanln(&reply)

		fmt.Println("Items \t Value")
		if reply == "no" {
			fmt.Println(index, "   \t  ", commission)
			fmt.Println("Hey there, your commission is :", commission)
			break
		}
	}
}
