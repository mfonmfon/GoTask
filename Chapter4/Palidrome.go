package main

import "fmt"

func main() {
	var reversed []int
	fmt.Print("Type a word ")
	fmt.Scanln(&reversed)

	if Palindrome(reversed) {
		fmt.Println("is palindrome")
	} else {
		fmt.Println("is not palindrome")
	}
}
func Palindrome(reversedTurned []int) bool {
	for index := 0; index < len(reversedTurned); index++ {
		if reversedTurned[index] != reversedTurned[index-1-index] {
			return false
		}
	}
	return true
}
