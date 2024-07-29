package main

import (
	"fmt"
)

var milesDrive float32
var gallonUsed float32
var response string

func main() {
	GasMillage()
}
func GasMillage() {
	for true {
		fmt.Print("Enter the miles driven: ")
		fmt.Scanln(&milesDrive)
		fmt.Print("Enter the gallons used: ")
		fmt.Scanln(&gallonUsed)

		fmt.Println("Do you want to continue? ")
		fmt.Scanln(&response)

		if response == "no" {
			var totalMiles float32 = milesDrive / gallonUsed
			fmt.Println("From the analysis, ", milesDrive, " Miles were driven and", gallonUsed, "gallons were used")
			fmt.Print(totalMiles, " miles were driven")
			break
		}
	}

}
