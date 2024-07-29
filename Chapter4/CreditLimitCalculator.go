package main

import (
	"errors"
	"fmt"
)

func main() {
	creditLimitCalculator()
}

var accountNumber int
var beginningBalance int
var itemsCharges int
var creditLimit int
var newBalance int

const allowedCreditLimit int = 500

func validateAccountNumber() {
	fmt.Println("Enter your account number: ")
	fmt.Scanln(&accountNumber)
	if accountNumber == 10 {
		creditLimitCalculator()
	} else {
		errors.New(`invalid Account Number: `)
	}
}

func creditLimitCalculator() {

	fmt.Println("Welcome to Credit Limit Calculator\nFollow the procedures below to check your credit limit")

	fmt.Println("Enter your beginning balance: ")
	fmt.Scanln(&beginningBalance)

	fmt.Println("Enter total of all items charged by the customer this month: ")
	fmt.Scanln(&itemsCharges)

	fmt.Println("Enter total of all credits applied to the customer’s account this month")
	fmt.Scanln(&creditLimit)

	if newBalance > allowedCreditLimit {
		newBalance = beginningBalance + itemsCharges - creditLimit
		fmt.Println(newBalance)
		fmt.Println("credit limit calculated")
	} else {
		fmt.Println("credit limit exceeded ")
	}

}
