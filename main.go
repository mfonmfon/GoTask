package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter your email: ")
	email, _ := reader.ReadString('\n')
	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	fmt.Print(" Hello ", name, " ", email, " ", password)
}
