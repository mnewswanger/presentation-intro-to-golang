package main

import "fmt"

func main() {
	printHello()
	print(" ")

	// Anonymous Functions
	func(s string) {
		fmt.Printf("%s!", s)
	}("world")

	println()
}

// Function to print Hello
func printHello() {
	fmt.Print("Hello")
}
