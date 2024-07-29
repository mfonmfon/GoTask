package main

import "fmt"

func main() {
	displayTables()
}
func displayTables() {
	fmt.Print("N \t ", "N1 \t ", "N2 \t ", "N3 \t ", "N4\t ")
	for index := 1; index < 4; index++ {
		fmt.Println(index, "  \t", index*2, "\n", index*3, " \t", index*4)
	}
}
