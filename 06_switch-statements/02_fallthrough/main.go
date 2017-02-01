package main

import "fmt"

func main(){
	switch "Marcus"{
	case "Tim":
		fmt.Println("holi T")
	case "Jenny":
		fmt.Println("holi J")
	case "Marcus":
		fmt.Println("holi M")
		fallthrough
	case "Medhi":
		fmt.Println("holi m")
		fallthrough
	case "Julian":
		fmt.Println("holi j")
	case "Sushant":
		fmt.Println("holi s")
	}
}