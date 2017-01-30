package main

import "fmt"

func zero(x *int){
	*x = 0
	//fmt.Println(x)
	//fmt.Println(&x)
}

func main(){
	x := 5
	//fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
}
