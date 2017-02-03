package main

import "fmt"

func main() {

	name := "Ricky"
	fmt.Println(name) // Ricky

	changeMe(name)

	fmt.Println(name) // Ricky
}

func changeMe(z string) {
	fmt.Println(z) // Ricky
	z = "Rocky"
	fmt.Println(z) // Rocky
}