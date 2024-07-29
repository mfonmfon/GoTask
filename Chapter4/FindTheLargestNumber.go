package main

import "fmt"

func main() {
	largestNumber()

}

var numberOfUnitSold [10]int
var largestNum int

func largestNumber() {

	for index := range numberOfUnitSold {
		fmt.Println("Enter the number of the things you sold: ")
		fmt.Scan(&numberOfUnitSold[index])

		if numberOfUnitSold[index] > largestNum {
			largestNum = numberOfUnitSold[index]
		}
	}
	fmt.Println("The largest number is", largestNum)
}
