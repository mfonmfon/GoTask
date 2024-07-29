package main

import "fmt"

func main() {
	taxCalculator()
}

var userEarning string

func taxCalculator() {
	for index := 0; index < 3; index++ {
		fmt.Print("Enter your Earning Price: ")
		fmt.Scan(&userEarning)

		if userEarning <= "30000" {
			fmt.Println("your tax rate is 15%")
		} else {
			fmt.Println("your tax rate is 20%")
		}
	}
}
