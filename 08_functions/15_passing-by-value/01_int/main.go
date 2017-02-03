package main

import "fmt"

func main(){
	age := 44
	changeMe(age)
	fmt.Println(age)
}
func changeMe(x int){
	fmt.Println(x)
	x = 24
}